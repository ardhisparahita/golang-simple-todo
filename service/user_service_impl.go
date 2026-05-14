package service

import (
	"context"
	"errors"
	"fmt"
	"golang-blog-api/domain"
	"golang-blog-api/exception"
	"golang-blog-api/model/web"
	"golang-blog-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) Register(ctx context.Context, req web.UserCreateRequest) (web.UserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return web.UserResponse{}, err
	}

	id := uuid.New()
	user := domain.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	result, err := s.Repo.Save(ctx, user)

	return web.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}, err
}

func (s *UserServiceImpl) Login(ctx context.Context, req web.UserLoginRequest) (web.UserLoginResponse, error) {
	user, err := s.Repo.FindByNameOrEmail(ctx, req.Identifier)

	if err != nil {
		return web.UserLoginResponse{}, errors.New("email/name or password wrong")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return web.UserLoginResponse{}, errors.New("email/name or password wrong")
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := viper.GetString("JWT_SECRET")
	fmt.Println(secretKey)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return web.UserLoginResponse{}, err
	}

	return web.UserLoginResponse{
		Token: tokenString,
	}, nil

}

func (s *UserServiceImpl) Update(ctx context.Context, id uuid.UUID, req web.UserUpdateRequest) (web.UserResponse, error) {
	user := domain.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	result, err := s.Repo.Update(ctx, id, user)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.UserResponse{}, &exception.NotFoundError{
				Message: "Not Found Error2",
			}
		}
	}

	return web.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}, err

}

func (s *UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.Repo.Delete(ctx, id)

	if err != nil {
		return &exception.NotFoundError{
			Message: "Not Found Error",
		}
	}

	return nil
}

func (s *UserServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (web.UserResponse, error) {
	user, err := s.Repo.FindByID(ctx, id)

	if err != nil {
		return web.UserResponse{}, &exception.NotFoundError{
			Message: "Not Found Error1",
		}
	}

	return web.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
func (s *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	users, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.UserResponse

	for _, user := range users {
		responses = append(responses, web.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return responses, nil
}
