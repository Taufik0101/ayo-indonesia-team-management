package service

import (
	"context"
	"errors"
	"fmt"
	"gin-ayo/database/models"
	"gin-ayo/dto"
	repositories "gin-ayo/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	DetailAccumulation struct {
		Team     *models.Team `json:"team"`
		TotalWin int64        `json:"total_win"`
	}

	DetailScheduleResponse struct {
		Schedule           *models.Schedule     `json:"schedule"`
		Status             string               `json:"status"`
		Score              string               `json:"score"`
		PlayerGoal         []*models.Player     `json:"player_goal"`
		DetailAccumulation []DetailAccumulation `json:"detail_accumulation"`
	}
)

type ScheduleService interface {
	Create(ctx context.Context, input dto.CreateSchedule) (*models.Schedule, error)
	Update(ctx context.Context, input dto.UpdateSchedule) (*models.Schedule, error)
	Delete(ctx context.Context, input dto.DeleteSchedule) (uuid.UUID, error)
	Detail(ctx context.Context, input dto.DetailSchedule) (*DetailScheduleResponse, error)
}

type scheduleService struct {
	scheduleRepository repositories.ScheduleRepositoryInterface
	resultRepository   repositories.ResultRepositoryInterface
	playerRepository   repositories.PlayerRepositoryInterface
}

func (t scheduleService) populateScore(input dto.DetailSchedule) string {
	output := ""

	currentResult, _ := t.resultRepository.FindOne(
		map[string]any{
			"schedule_id": input.ID,
		}, nil, nil)

	if currentResult != nil {
		output = fmt.Sprintf("%d:%d", currentResult.ScoreHome, currentResult.ScoreAway)
	}

	return output
}

func (t scheduleService) populateStatus(schedule *models.Schedule) string {
	output := "Draw"

	if schedule.WinnerTeamID != uuid.Nil {
		output = fmt.Sprintf("%s Menang", schedule.WinnerTeam.Name)
	}

	return output
}

func (t scheduleService) populatePlayerGoal(input dto.DetailSchedule) []*models.Player {
	output := make([]*models.Player, 0)

	currentResult, _ := t.resultRepository.FindOne(
		map[string]any{
			"schedule_id": input.ID,
		}, nil, []string{"DetailResult"})

	if currentResult != nil {
		mapPlayerGoal := make(map[uuid.UUID]int64)

		if len(currentResult.DetailResult) > 0 {
			arrPlayerID := make([]uuid.UUID, 0)

			for _, dR := range currentResult.DetailResult {
				if _, exists := mapPlayerGoal[dR.PlayerID]; exists {
					temp := mapPlayerGoal[dR.PlayerID]
					temp++
					mapPlayerGoal[dR.PlayerID] = temp
				} else {
					temp := mapPlayerGoal[dR.PlayerID]
					temp = 1
					mapPlayerGoal[dR.PlayerID] = temp
					arrPlayerID = append(arrPlayerID, dR.PlayerID)
				}
			}

			maxVal := int64(0)
			for _, v := range mapPlayerGoal {
				if v > maxVal {
					maxVal = v
				}
			}

			topKeys := make([]uuid.UUID, 0)
			for k, v := range mapPlayerGoal {
				if v == maxVal {
					topKeys = append(topKeys, k)
				}
			}

			findPlayer, _ := t.playerRepository.FindMany(
				map[string]any{
					"id in (?)": topKeys,
				}, nil, nil)

			if findPlayer != nil {
				output = findPlayer
			}
		}
	}

	return output
}

func (t scheduleService) populateDetailAccumulation(schedule *models.Schedule) []DetailAccumulation {
	output := make([]DetailAccumulation, 0)

	findScheduleHomeWin, _ := t.scheduleRepository.FindMany(
		map[string]any{
			"winner_team_id": schedule.HomeTeamID,
			"date <= ?":      schedule.Date,
		}, nil, nil)

	output = append(output, DetailAccumulation{
		Team:     schedule.HomeTeam,
		TotalWin: int64(len(findScheduleHomeWin)),
	})

	findScheduleAwayWin, _ := t.scheduleRepository.FindMany(
		map[string]any{
			"winner_team_id": schedule.AwayTeamID,
			"date <= ?":      schedule.Date,
		}, nil, nil)

	output = append(output, DetailAccumulation{
		Team:     schedule.AwayTeam,
		TotalWin: int64(len(findScheduleAwayWin)),
	})

	return output
}

func (t scheduleService) Detail(ctx context.Context, input dto.DetailSchedule) (*DetailScheduleResponse, error) {
	output := new(DetailScheduleResponse)

	currentSchedule, err := t.scheduleRepository.FindOne(
		map[string]any{
			"id": input.ID,
		}, nil, []string{"HomeTeam", "AwayTeam", "WinnerTeam"})

	if err != nil {
		return nil, err
	}

	output.Schedule = currentSchedule
	output.Score = t.populateScore(input)
	output.Status = t.populateStatus(currentSchedule)
	output.PlayerGoal = t.populatePlayerGoal(input)
	output.DetailAccumulation = t.populateDetailAccumulation(currentSchedule)

	return output, nil
}

func (t scheduleService) Create(ctx context.Context, input dto.CreateSchedule) (*models.Schedule, error) {
	output := new(models.Schedule)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	localTime, err := time.ParseInLocation("2006-01-02 15:04", input.DateTime, loc)
	if err != nil {
		return nil, errors.New("invalid time format")
	}

	if !localTime.After(time.Now().In(loc)) {
		return nil, errors.New("datetime must be greater than current time (Jakarta)")
	}

	if input.HomeTeamID == input.AwayTeamID {
		return nil, errors.New("home team and away team must be different")
	}

	tUTC := localTime.UTC()

	findSchedule, _ := t.scheduleRepository.FindOne(map[string]any{
		"date": tUTC,
	}, nil, nil)

	if findSchedule != nil {
		return nil, errors.New("schedule already exists")
	}

	scheduleList := make([]*models.Schedule, 0)
	scheduleList = append(scheduleList, &models.Schedule{
		Date:       tUTC,
		HomeTeamID: uuid.MustParse(input.HomeTeamID),
		AwayTeamID: uuid.MustParse(input.AwayTeamID),
	})

	errTx := t.scheduleRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		outputs, err := t.scheduleRepository.Create(ctx, scheduleList)
		if err != nil {
			logrus.Errorf("[scheduleService][Create] failed to create schedule, error: %v", err)
			return err
		}

		output = outputs[0]

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[scheduleService][Create] failed to commit transaction, error: %v", errTx)
		return nil, errTx
	}

	return output, nil
}

func (t scheduleService) Update(ctx context.Context, input dto.UpdateSchedule) (*models.Schedule, error) {
	output := new(models.Schedule)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	localTime, err := time.ParseInLocation("2006-01-02 15:04", input.DateTime, loc)
	if err != nil {
		return nil, errors.New("invalid time format")
	}

	if !localTime.After(time.Now().In(loc)) {
		return nil, errors.New("datetime must be greater than current time (Jakarta)")
	}

	if input.HomeTeamID == input.AwayTeamID {
		return nil, errors.New("home team and away team must be different")
	}

	tUTC := localTime.UTC()

	findSchedule, _ := t.scheduleRepository.FindOne(map[string]any{
		"date": tUTC,
	}, map[string]any{
		"id": input.ID,
	}, nil)

	if findSchedule != nil {
		return nil, errors.New("schedule already exists")
	}

	scheduleList := make([]*models.Schedule, 0)
	scheduleList = append(scheduleList, &models.Schedule{
		BaseModel: models.BaseModel{
			ID: uuid.MustParse(input.ID),
		},
		Date:       tUTC,
		HomeTeamID: uuid.MustParse(input.HomeTeamID),
		AwayTeamID: uuid.MustParse(input.AwayTeamID),
	})

	errTx := t.scheduleRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		outputs, err := t.scheduleRepository.Update(ctx, scheduleList, nil)
		if err != nil {
			logrus.Errorf("[scheduleService][Create] failed to update schedule, error: %v", err)
			return err
		}

		output = outputs[0]

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[scheduleService][Create] failed to commit transaction, error: %v", errTx)
		return nil, errTx
	}

	return output, nil
}

func (t scheduleService) Delete(ctx context.Context, input dto.DeleteSchedule) (uuid.UUID, error) {
	scheduleList := make([]uuid.UUID, 0)
	scheduleList = append(scheduleList, uuid.MustParse(input.ID))

	errTx := t.scheduleRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		err := t.scheduleRepository.Delete(ctx, scheduleList)
		if err != nil {
			logrus.Errorf("[scheduleService][Create] failed to delete schedule, error: %v", err)
			return err
		}

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[scheduleService][Create] failed to commit transaction, error: %v", errTx)
		return uuid.Nil, errTx
	}

	return uuid.MustParse(input.ID), nil
}

func NewScheduleService(
	scheduleRepository repositories.ScheduleRepositoryInterface,
	resultRepository repositories.ResultRepositoryInterface,
	playerRepository repositories.PlayerRepositoryInterface,
) ScheduleService {
	return &scheduleService{scheduleRepository: scheduleRepository, resultRepository: resultRepository, playerRepository: playerRepository}
}
