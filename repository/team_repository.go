package repositories

import (
	"context"
	"gin-ayo/database/models"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TeamRepositoryInterface interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Team, error)
	Create(ctx context.Context, teams []*models.Team) ([]*models.Team, error)
	Update(ctx context.Context, teams []*models.Team, omitRelations []string) ([]*models.Team, error)
	Delete(ctx context.Context, uuids []uuid.UUID) error
}

type teamRepository struct {
	db *gorm.DB
}

func (u teamRepository) Create(ctx context.Context, teams []*models.Team) ([]*models.Team, error) {
	if len(teams) < 1 {
		return make([]*models.Team, 0), nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Save(&teams).Error
	if err != nil {
		log.Errorf("[TeamRepository][Create] failed to exec Save, error: %v", err)
	}
	return teams, err
}

func (u teamRepository) Update(ctx context.Context, teams []*models.Team, omitRelations []string) ([]*models.Team, error) {
	if len(teams) < 1 {
		return nil, nil
	}

	db := ExtractTx(ctx, u.db)

	if len(omitRelations) > 0 {
		db = db.Omit(omitRelations...)
	}

	err := db.Save(&teams).Error
	if err != nil {
		log.Errorf("[teamRepository][Update] failed to exec Save, error: %v", err)
		return nil, err
	}

	return teams, nil
}

func (u teamRepository) Delete(ctx context.Context, uuids []uuid.UUID) error {
	if len(uuids) < 1 {
		return nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Where("id IN ?", uuids).Delete(&models.Team{}).Error
	if err != nil {
		log.Errorf("[teamRepository][Delete] failed to exec Delete, error: %v", err)
	}
	return err
}

func (u teamRepository) FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Team, error) {
	var Team *models.Team

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

	err := query.First(&Team).Error
	if err != nil {
		log.Errorf("[TeamRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return Team, nil
}

func (u teamRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
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

func NewTeamRepository(db *gorm.DB) TeamRepositoryInterface {
	return &teamRepository{db: db}
}
