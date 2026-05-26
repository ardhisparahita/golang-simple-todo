package service

import (
	"context"
	"golang-blog-api/model/web"

	"github.com/google/uuid"
)

type TodoService interface {
	Create(ctx context.Context, userId uuid.UUID, req web.TodoCreateRequest) (web.TodoResponse, error)
	Update(ctx context.Context, id uuid.UUID, userId uuid.UUID, req web.TodoUpdateRequest) (web.TodoResponse, error)
	Delete(ctx context.Context, id uuid.UUID, userId uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID, userId uuid.UUID) (web.TodoResponse, error)
	FindByUserID(ctx context.Context, userId uuid.UUID) ([]web.TodoResponse, error)
}
