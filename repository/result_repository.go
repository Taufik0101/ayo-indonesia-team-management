package repositories

import (
	"context"
	"gin-ayo/database/models"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ResultRepositoryInterface interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Result, error)
	Create(ctx context.Context, results []*models.Result) ([]*models.Result, error)
	Update(ctx context.Context, results []*models.Result, omitRelations []string) ([]*models.Result, error)
	Delete(ctx context.Context, uuids []uuid.UUID) error
}

type resultRepository struct {
	db *gorm.DB
}

func (u resultRepository) Create(ctx context.Context, results []*models.Result) ([]*models.Result, error) {
	if len(results) < 1 {
		return make([]*models.Result, 0), nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Save(&results).Error
	if err != nil {
		log.Errorf("[ResultRepository][Create] failed to exec Save, error: %v", err)
	}
	return results, err
}

func (u resultRepository) Update(ctx context.Context, results []*models.Result, omitRelations []string) ([]*models.Result, error) {
	if len(results) < 1 {
		return nil, nil
	}

	db := ExtractTx(ctx, u.db)

	if len(omitRelations) > 0 {
		db = db.Omit(omitRelations...)
	}

	err := db.Save(&results).Error
	if err != nil {
		log.Errorf("[resultRepository][Update] failed to exec Save, error: %v", err)
		return nil, err
	}

	return results, nil
}

func (u resultRepository) Delete(ctx context.Context, uuids []uuid.UUID) error {
	if len(uuids) < 1 {
		return nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Where("id IN ?", uuids).Delete(&models.Result{}).Error
	if err != nil {
		log.Errorf("[resultRepository][Delete] failed to exec Delete, error: %v", err)
	}
	return err
}

func (u resultRepository) FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Result, error) {
	var Result *models.Result

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

	err := query.First(&Result).Error
	if err != nil {
		log.Errorf("[ResultRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return Result, nil
}

func (u resultRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
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

func NewResultRepository(db *gorm.DB) ResultRepositoryInterface {
	return &resultRepository{db: db}
}
