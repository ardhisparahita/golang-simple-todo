package repository

import (
	"context"
	"golang-blog-api/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) Save(ctx context.Context, user domain.User) (domain.User, error) {
	err := r.DB.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(ctx context.Context, id uuid.UUID, user domain.User) (domain.User, error) {
	result := r.DB.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Updates(user)

	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).Delete(&domain.User{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var user domain.User

	err := r.DB.WithContext(ctx).First(&user, "id = ?", id).Error
	return user, err
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	err := r.DB.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindByNameOrEmail(ctx context.Context, identifier string) (domain.User, error) {
	var user domain.User

	err := r.DB.WithContext(ctx).Where("email = ? OR name = ? ", identifier, identifier).First(&user).Error

	return user, err
}
