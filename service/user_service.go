package service

import (
	"context"
	"golang-blog-api/model/web"

	"github.com/google/uuid"
)

type UserService interface {
	Register(ctx context.Context, req web.UserCreateRequest) (web.UserResponse, error)
	Login(ctx context.Context, req web.UserLoginRequest) (web.UserLoginResponse, error)
	Update(ctx context.Context, id uuid.UUID, req web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (web.UserResponse, error)
	FindAll(ctx context.Context) ([]web.UserResponse, error)
}
