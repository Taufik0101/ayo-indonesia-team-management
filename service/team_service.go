package service

import (
	"context"
	"gin-ayo/database/models"
	"gin-ayo/dto"
	repositories "gin-ayo/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TeamService interface {
	Create(ctx context.Context, input dto.CreateTeam) (*models.Team, error)
	Update(ctx context.Context, input dto.UpdateTeam) (*models.Team, error)
	Delete(ctx context.Context, input dto.DeleteTeam) (uuid.UUID, error)
}

type teamService struct {
	teamRepository repositories.TeamRepositoryInterface
}

func (t teamService) Create(ctx context.Context, input dto.CreateTeam) (*models.Team, error) {
	output := new(models.Team)
	teamList := make([]*models.Team, 0)
	teamList = append(teamList, &models.Team{
		Name:          input.Name,
		Logo:          input.Logo,
		Address:       input.Address,
		Year:          input.Year,
		ProvinceID:    uuid.MustParse(input.ProvinceID),
		DistrictID:    uuid.MustParse(input.DistrictID),
		SubDistrictID: uuid.MustParse(input.SubDistrictID),
		VillageID:     uuid.MustParse(input.VillageID),
	})

	errTx := t.teamRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		outputs, err := t.teamRepository.Create(ctx, teamList)
		if err != nil {
			logrus.Errorf("[teamService][Create] failed to create team, error: %v", err)
			return err
		}

		output = outputs[0]

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[teamService][Create] failed to commit transaction, error: %v", errTx)
		return nil, errTx
	}

	return output, nil
}

func (t teamService) Update(ctx context.Context, input dto.UpdateTeam) (*models.Team, error) {
	output := new(models.Team)
	teamList := make([]*models.Team, 0)
	teamList = append(teamList, &models.Team{
		BaseModel: models.BaseModel{
			ID: uuid.MustParse(input.ID),
		},
		Name:          input.Name,
		Logo:          input.Logo,
		Address:       input.Address,
		Year:          input.Year,
		ProvinceID:    uuid.MustParse(input.ProvinceID),
		DistrictID:    uuid.MustParse(input.DistrictID),
		SubDistrictID: uuid.MustParse(input.SubDistrictID),
		VillageID:     uuid.MustParse(input.VillageID),
	})

	errTx := t.teamRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		outputs, err := t.teamRepository.Update(ctx, teamList, nil)
		if err != nil {
			logrus.Errorf("[teamService][Create] failed to update team, error: %v", err)
			return err
		}

		output = outputs[0]

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[teamService][Create] failed to commit transaction, error: %v", errTx)
		return nil, errTx
	}

	return output, nil
}

func (t teamService) Delete(ctx context.Context, input dto.DeleteTeam) (uuid.UUID, error) {
	teamList := make([]uuid.UUID, 0)
	teamList = append(teamList, uuid.MustParse(input.ID))

	errTx := t.teamRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		err := t.teamRepository.Delete(ctx, teamList)
		if err != nil {
			logrus.Errorf("[teamService][Create] failed to delete team, error: %v", err)
			return err
		}

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[teamService][Create] failed to commit transaction, error: %v", errTx)
		return uuid.Nil, errTx
	}

	return uuid.MustParse(input.ID), nil
}

func NewTeamService(teamRepository repositories.TeamRepositoryInterface) TeamService {
	return &teamService{teamRepository: teamRepository}
}
