package repository

import (
	"context"
	"golang-blog-api/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, id uuid.UUID, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByNameOrEmail(ctx context.Context, identifier string) (domain.User, error)
}
