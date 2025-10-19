package repositories

import (
	"context"
	"gin-ayo/database/models"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DetailResultRepositoryInterface interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.DetailResult, error)
	Create(ctx context.Context, detailResults []*models.DetailResult) ([]*models.DetailResult, error)
	Update(ctx context.Context, detailResults []*models.DetailResult, omitRelations []string) ([]*models.DetailResult, error)
	Delete(ctx context.Context, uuids []uuid.UUID) error
}

type detailResultRepository struct {
	db *gorm.DB
}

func (u detailResultRepository) Create(ctx context.Context, detailResults []*models.DetailResult) ([]*models.DetailResult, error) {
	if len(detailResults) < 1 {
		return make([]*models.DetailResult, 0), nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Save(&detailResults).Error
	if err != nil {
		log.Errorf("[DetailResultRepository][Create] failed to exec Save, error: %v", err)
	}
	return detailResults, err
}

func (u detailResultRepository) Update(ctx context.Context, detailResults []*models.DetailResult, omitRelations []string) ([]*models.DetailResult, error) {
	if len(detailResults) < 1 {
		return nil, nil
	}

	db := ExtractTx(ctx, u.db)

	if len(omitRelations) > 0 {
		db = db.Omit(omitRelations...)
	}

	err := db.Save(&detailResults).Error
	if err != nil {
		log.Errorf("[detailResultRepository][Update] failed to exec Save, error: %v", err)
		return nil, err
	}

	return detailResults, nil
}

func (u detailResultRepository) Delete(ctx context.Context, uuids []uuid.UUID) error {
	if len(uuids) < 1 {
		return nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Where("id IN ?", uuids).Delete(&models.DetailResult{}).Error
	if err != nil {
		log.Errorf("[detailResultRepository][Delete] failed to exec Delete, error: %v", err)
	}
	return err
}

func (u detailResultRepository) FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.DetailResult, error) {
	var DetailResult *models.DetailResult

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

	err := query.First(&DetailResult).Error
	if err != nil {
		log.Errorf("[DetailResultRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return DetailResult, nil
}

func (u detailResultRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
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

func NewDetailResultRepository(db *gorm.DB) DetailResultRepositoryInterface {
	return &detailResultRepository{db: db}
}
