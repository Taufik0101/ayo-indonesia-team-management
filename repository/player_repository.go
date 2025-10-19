package repositories

import (
	"context"
	"gin-ayo/database/models"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PlayerRepositoryInterface interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Player, error)
	FindMany(whereClause map[string]any, whereNotClause map[string]any, relations []string) ([]*models.Player, error)
	Create(ctx context.Context, players []*models.Player) ([]*models.Player, error)
	Update(ctx context.Context, players []*models.Player, omitRelations []string) ([]*models.Player, error)
	Delete(ctx context.Context, uuids []uuid.UUID) error
}

type playerRepository struct {
	db *gorm.DB
}

func (u playerRepository) FindMany(whereClause map[string]any, whereNotClause map[string]any, relations []string) ([]*models.Player, error) {
	var Player []*models.Player

	query := u.db

	for queryWhere, value := range whereClause {
		query = query.Where(queryWhere, value)
	}

	for queryNotWhere, value := range whereNotClause {
		query = query.Not(queryNotWhere, value)
	}

	for _, relation := range relations {
		query = query.Preload(relation)
	}

	err := query.Find(&Player).Error
	if err != nil {
		log.Errorf("[PlayerRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return Player, nil
}

func (u playerRepository) Create(ctx context.Context, players []*models.Player) ([]*models.Player, error) {
	if len(players) < 1 {
		return make([]*models.Player, 0), nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Save(&players).Error
	if err != nil {
		log.Errorf("[PlayerRepository][Create] failed to exec Save, error: %v", err)
	}
	return players, err
}

func (u playerRepository) Update(ctx context.Context, players []*models.Player, omitRelations []string) ([]*models.Player, error) {
	if len(players) < 1 {
		return nil, nil
	}

	db := ExtractTx(ctx, u.db)

	if len(omitRelations) > 0 {
		db = db.Omit(omitRelations...)
	}

	err := db.Save(&players).Error
	if err != nil {
		log.Errorf("[playerRepository][Update] failed to exec Save, error: %v", err)
		return nil, err
	}

	return players, nil
}

func (u playerRepository) Delete(ctx context.Context, uuids []uuid.UUID) error {
	if len(uuids) < 1 {
		return nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Where("id IN ?", uuids).Delete(&models.Player{}).Error
	if err != nil {
		log.Errorf("[playerRepository][Delete] failed to exec Delete, error: %v", err)
	}
	return err
}

func (u playerRepository) FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Player, error) {
	var Player *models.Player

	query := u.db

	for queryWhere, value := range whereClause {
		query = query.Where(queryWhere, value)
	}

	for queryNotWhere, value := range whereNotClause {
		query = query.Not(queryNotWhere, value)
	}

	for _, relation := range relations {
		query = query.Preload(relation)
	}

	err := query.First(&Player).Error
	if err != nil {
		log.Errorf("[PlayerRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return Player, nil
}

func (u playerRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx := u.db.WithContext(ctx).Begin()

	if tx.Error != nil {
		log.Errorf("failed to begin transaction: %v", tx.Error)
		return tx.Error
	}

	if err := tFunc(InjectTx(ctx, tx)); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Errorf("failed to commit transaction: %v", tx.Error)
		return err
	}

	return nil
}

func NewPlayerRepository(db *gorm.DB) PlayerRepositoryInterface {
	return &playerRepository{db: db}
}
