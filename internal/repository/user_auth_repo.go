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
