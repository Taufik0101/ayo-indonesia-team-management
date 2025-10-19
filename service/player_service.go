package service

import (
	"context"
	"errors"
	"gin-ayo/database/models"
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	repositories "gin-ayo/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PlayerService interface {
	Create(ctx context.Context, input dto.CreatePlayer) (*models.Player, error)
	Update(ctx context.Context, input dto.UpdatePlayer) (*models.Player, error)
	Delete(ctx context.Context, input dto.DeletePlayer) (uuid.UUID, error)
}

type playerService struct {
	playerRepository repositories.PlayerRepositoryInterface
}

func (t playerService) Create(ctx context.Context, input dto.CreatePlayer) (*models.Player, error) {
	output := new(models.Player)

	findPlayer, _ := t.playerRepository.FindOne(map[string]any{
		"number":  input.Number,
		"team_id": input.TeamID,
	}, nil, nil)

	if findPlayer != nil {
		return nil, errors.New("player number already exists")
	}

	playerList := make([]*models.Player, 0)
	playerList = append(playerList, &models.Player{
		Name:     input.Name,
		Weight:   input.Weight,
		Height:   input.Height,
		Number:   input.Number,
		Position: utils.PlayerPositionType(input.Position),
		TeamID:   uuid.MustParse(input.TeamID),
	})

	errTx := t.playerRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		outputs, err := t.playerRepository.Create(ctx, playerList)
		if err != nil {
			logrus.Errorf("[playerService][Create] failed to create player, error: %v", err)
			return err
		}

		output = outputs[0]

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[playerService][Create] failed to commit transaction, error: %v", errTx)
		return nil, errTx
	}

	return output, nil
}

func (t playerService) Update(ctx context.Context, input dto.UpdatePlayer) (*models.Player, error) {
	output := new(models.Player)

	findPlayer, _ := t.playerRepository.FindOne(map[string]any{
		"number":  input.Number,
		"team_id": input.TeamID,
	}, map[string]any{
		"id": input.ID,
	}, nil)

	if findPlayer != nil {
		return nil, errors.New("player number already exists")
	}

	playerList := make([]*models.Player, 0)
	playerList = append(playerList, &models.Player{
		BaseModel: models.BaseModel{
			ID: uuid.MustParse(input.ID),
		},
		Name:     input.Name,
		Weight:   input.Weight,
		Height:   input.Height,
		Number:   input.Number,
		Position: utils.PlayerPositionType(input.Position),
		TeamID:   uuid.MustParse(input.TeamID),
	})

	errTx := t.playerRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		outputs, err := t.playerRepository.Update(ctx, playerList, nil)
		if err != nil {
			logrus.Errorf("[playerService][Create] failed to update player, error: %v", err)
			return err
		}

		output = outputs[0]

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[playerService][Create] failed to commit transaction, error: %v", errTx)
		return nil, errTx
	}

	return output, nil
}

func (t playerService) Delete(ctx context.Context, input dto.DeletePlayer) (uuid.UUID, error) {
	playerList := make([]uuid.UUID, 0)
	playerList = append(playerList, uuid.MustParse(input.ID))

	errTx := t.playerRepository.WithinTransaction(ctx, func(ctx context.Context) error {
		err := t.playerRepository.Delete(ctx, playerList)
		if err != nil {
			logrus.Errorf("[playerService][Create] failed to delete player, error: %v", err)
			return err
		}

		return nil
	})

	if errTx != nil {
		logrus.Errorf("[playerService][Create] failed to commit transaction, error: %v", errTx)
		return uuid.Nil, errTx
	}

	return uuid.MustParse(input.ID), nil
}

func NewPlayerService(playerRepository repositories.PlayerRepositoryInterface) PlayerService {
	return &playerService{playerRepository: playerRepository}
}
