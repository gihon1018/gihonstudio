package repository

import (
	"context"
	. "gihonstudio/internal/model"

	"gorm.io/gorm"
)

type UserAuthRepository struct {
	db *gorm.DB
}

func NewUserAuthRepository(db *gorm.DB) *UserAuthRepository {
	return &UserAuthRepository{db: db}
}

func (r *UserAuthRepository) Create(ctx context.Context, userAuth *UserAuth) error {
	return gorm.G[UserAuth](r.db).Create(ctx, userAuth)
}

func (r *UserAuthRepository) CreateWithRows(ctx context.Context, userAuth *UserAuth) (int64, error) {
	result := gorm.WithResult()
	return result.RowsAffected, gorm.G[UserAuth](r.db, result).Create(ctx, userAuth)
}

func (r *UserAuthRepository) CreateOmit(ctx context.Context, userAuth *UserAuth, selectStr []string) error {
	return gorm.G[UserAuth](r.db).Omit(selectStr...).Create(ctx, userAuth)
}

func (r *UserAuthRepository) CreateInBatches(ctx context.Context, userAuths *[]UserAuth, batchSize int) error {
	return gorm.G[UserAuth](r.db).CreateInBatches(ctx, userAuths, batchSize)
}
