package repository

import (
	"context"
	"golang-blog-api/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{
		DB: db,
	}
}

func (r *TodoRepositoryImpl) Save(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	err := r.DB.WithContext(ctx).Create(&todo).Error
	return todo, err
}

func (r *TodoRepositoryImpl) Update(ctx context.Context, id uuid.UUID, todo domain.Todo) (domain.Todo, error) {
	result := r.DB.WithContext(ctx).Model(&domain.Todo{}).Where("id = ? ", id).Updates(map[string]any{
		"title":       todo.Title,
		"description": todo.Description,
		"completed":   todo.Completed,
	})

	if result.Error != nil {
		return todo, result.Error
	}

	if result.RowsAffected == 0 {
		return todo, gorm.ErrRecordNotFound
	}

	return todo, nil
}

func (r *TodoRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).Delete(&domain.Todo{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
func (r *TodoRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Todo, error) {
	var todo domain.Todo

	err := r.DB.WithContext(ctx).First(&todo, "id = ?", id).Error
	return todo, err
}
func (r *TodoRepositoryImpl) FindByUserID(ctx context.Context, userId uuid.UUID) ([]domain.Todo, error) {
	var todos []domain.Todo

	err := r.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&todos).Error
	return todos, err
}
