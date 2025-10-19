package service

import (
	"context"
	"gin-ayo/database/models"
	"gin-ayo/dto"
	repositories "gin-ayo/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ResultService interface {
	Create(ctx context.Context, input dto.CreateResult) (*models.Result, error)
}

type resultService struct {
	resultRepository       repositories.ResultRepositoryInterface
	detailResultRepository repositories.DetailResultRepositoryInterface
	scheduleRepository     repositories.ScheduleRepositoryInterface
}

func (t resultService) Create(ctx context.Context, input dto.CreateResult) (*models.Result, error) {
	output := new(models.Result)

	currentSchedule, err := t.scheduleRepository.FindOne(
		map[string]any{
			"id": input.ScheduleID,
		}, nil, nil)

	if err != nil {
		return nil, err
	}

	var currentResult = &models.Result{
		ScoreHome:  input.ScoreHome,
		ScoreAway:  input.ScoreAway,
		ScheduleID: uuid.MustParse(input.ScheduleID),
	}

	currentSchedule.IsFinished = true
	if input.ScoreHome > input.ScoreAway {
		currentSchedule.WinnerTeamID = currentSchedule.HomeTeamID
		currentResult.WinnerTeamID = currentSchedule.HomeTeamID
	}
	if input.ScoreHome < input.ScoreAway {
		currentSchedule.WinnerTeamID = currentSchedule.AwayTeamID
		currentResult.WinnerTeamID = currentSchedule.AwayTeamID
	}

	scheduleList := make([]*models.Schedule, 0)
	scheduleList = append(scheduleList, currentSchedule)

	resultList := make([]*models.Result, 0)
	resultList = append(resultList, currentResult)

	detailResultList := make([]*models.DetailResult, 0)

	errTx := t.resultRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, res := range resultList {
			res.ID = uuid.New()

			for _, dr := range input.Details {
				detailResultList = append(detailResultList, &models.DetailResult{
					GoalTime:  dr.GoalTime,
					ResultID:  res.ID,
					IsPenalty: dr.IsPenalty,
					PlayerID:  uuid.MustParse(dr.PlayerID),
				})
			}
		}

		outputs, err := t.resultRepository.Create(ctx, resultList)
		if err != nil {
			logrus.Errorf("[resultService][Create] failed to create result, error: %v", err)
			return err
		}

		_, err = t.scheduleRepository.Update(ctx, scheduleList, nil)
		if err != nil {
			logrus.Errorf("[resultService][Create] failed to create result, error: %v", err)
			return err
		}

		outputDetail, err := t.detailResultRepository.Create(ctx, detailResultList)
		if err != nil {
			logrus.Errorf("[resultService][Create] failed to create detail result, error: %v", err)
			return err
		}

		output = outputs[0]
		output.DetailResult = outputDetail

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[resultService][Create] failed to commit transaction, error: %v", errTx)
		return nil, errTx
	}

	return output, nil
}

func NewResultService(
	resultRepository repositories.ResultRepositoryInterface,
	detailResultRepository repositories.DetailResultRepositoryInterface,
	scheduleRepository repositories.ScheduleRepositoryInterface,
) ResultService {
	return &resultService{resultRepository: resultRepository, detailResultRepository: detailResultRepository, scheduleRepository: scheduleRepository}
}
