package repositories

import (
	"context"
	"gin-ayo/database/models"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ScheduleRepositoryInterface interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Schedule, error)
	FindMany(whereClause map[string]any, whereNotClause map[string]any, relations []string) ([]*models.Schedule, error)
	Create(ctx context.Context, schedules []*models.Schedule) ([]*models.Schedule, error)
	Update(ctx context.Context, schedules []*models.Schedule, omitRelations []string) ([]*models.Schedule, error)
	Delete(ctx context.Context, uuids []uuid.UUID) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func (u scheduleRepository) FindMany(whereClause map[string]any, whereNotClause map[string]any, relations []string) ([]*models.Schedule, error) {
	var Schedule []*models.Schedule

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

	err := query.Find(&Schedule).Error
	if err != nil {
		log.Errorf("[ScheduleRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return Schedule, nil
}

func (u scheduleRepository) Create(ctx context.Context, schedules []*models.Schedule) ([]*models.Schedule, error) {
	if len(schedules) < 1 {
		return make([]*models.Schedule, 0), nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Save(&schedules).Error
	if err != nil {
		log.Errorf("[ScheduleRepository][Create] failed to exec Save, error: %v", err)
	}
	return schedules, err
}

func (u scheduleRepository) Update(ctx context.Context, schedules []*models.Schedule, omitRelations []string) ([]*models.Schedule, error) {
	if len(schedules) < 1 {
		return nil, nil
	}

	db := ExtractTx(ctx, u.db)

	if len(omitRelations) > 0 {
		db = db.Omit(omitRelations...)
	}

	err := db.Save(&schedules).Error
	if err != nil {
		log.Errorf("[scheduleRepository][Update] failed to exec Save, error: %v", err)
		return nil, err
	}

	return schedules, nil
}

func (u scheduleRepository) Delete(ctx context.Context, uuids []uuid.UUID) error {
	if len(uuids) < 1 {
		return nil
	}

	db := ExtractTx(ctx, u.db)

	err := db.Where("id IN ?", uuids).Delete(&models.Schedule{}).Error
	if err != nil {
		log.Errorf("[scheduleRepository][Delete] failed to exec Delete, error: %v", err)
	}
	return err
}

func (u scheduleRepository) FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.Schedule, error) {
	var Schedule *models.Schedule

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

	err := query.First(&Schedule).Error
	if err != nil {
		log.Errorf("[ScheduleRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return Schedule, nil
}

func (u scheduleRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
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

func NewScheduleRepository(db *gorm.DB) ScheduleRepositoryInterface {
	return &scheduleRepository{db: db}
}
