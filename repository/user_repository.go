package repositories

import (
	"context"
	"gin-ayo/database/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) FindOne(whereClause map[string]any, whereNotClause map[string]any, relations []string) (*models.User, error) {
	var User *models.User

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

	err := query.First(&User).Error
	if err != nil {
		log.Errorf("[UserRepository][FindOne] failed to exec First, error: %v", err)
		return nil, err
	}

	return User, nil
}

func (u userRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
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

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: db}
}
