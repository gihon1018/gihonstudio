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

func (r *UserAuthRepository) CreateOmit(ctx context.Context, userAuth *UserAuth, omitStrs []string) error {
	return gorm.G[UserAuth](r.db).Omit(omitStrs...).Create(ctx, userAuth)
}

func (r *UserAuthRepository) CreateInBatches(ctx context.Context, userAuths *[]UserAuth, batchSize int) error {
	return gorm.G[UserAuth](r.db).CreateInBatches(ctx, userAuths, batchSize)
}

func (r *UserAuthRepository) FirstByUid(ctx context.Context, uid uint64) (UserAuth, error) {
	return gorm.G[UserAuth](r.db).Where("uid = ?", uid).First(ctx)
}

func (r *UserAuthRepository) FirstByUsername(ctx context.Context, username string) (UserAuth, error) {
	return gorm.G[UserAuth](r.db).Where("username = ?", username).First(ctx)
}

func (r *UserAuthRepository) UpdateByUid(ctx context.Context, uid uint64, column string, value string) (int, error) {
	return gorm.G[UserAuth](r.db).Where("uid = ?", uid).Update(ctx, column, value)
}

func (r *UserAuthRepository) UpdateByUsername(ctx context.Context, username string, column string, value string) (int, error) {
	return gorm.G[UserAuth](r.db).Where("username = ?", username).Update(ctx, column, value)
}

func (r *UserAuthRepository) Updates(ctx context.Context, userAuth UserAuth) (int, error) {
	return gorm.G[UserAuth](r.db).Where("uid = ?", userAuth.Uid).Updates(ctx, userAuth)
}
