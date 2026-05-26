package repository

import (
	"context"
	"golang-blog-api/domain"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Save(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	Update(ctx context.Context, id uuid.UUID, todo domain.Todo) (domain.Todo, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (domain.Todo, error)
	FindByUserID(ctx context.Context, userId uuid.UUID) ([]domain.Todo, error)
}
