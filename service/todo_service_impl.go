package service

import (
	"context"
	"errors"
	"golang-blog-api/domain"
	"golang-blog-api/helper"
	"golang-blog-api/model/web"
	"golang-blog-api/repository"

	"github.com/google/uuid"
)

type TodoServiceImpl struct {
	Repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &TodoServiceImpl{Repo: repo}
}

func (s *TodoServiceImpl) Create(ctx context.Context, userId uuid.UUID, req web.TodoCreateRequest) (web.TodoResponse, error) {
	todo := domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		UserID:      userId,
	}

	result, err := s.Repo.Save(ctx, todo)
	if err != nil {
		return web.TodoResponse{}, err
	}

	return helper.ToTodoResponse(result), nil
}
func (s *TodoServiceImpl) Update(
	ctx context.Context,
	id uuid.UUID,
	userId uuid.UUID,
	req web.TodoUpdateRequest,
) (web.TodoResponse, error) {

	existingTodo, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.TodoResponse{}, err
	}

	if existingTodo.UserID != userId {
		return web.TodoResponse{}, errors.New("Forbidden")
	}

	existingTodo.Title = req.Title
	existingTodo.Description = req.Description
	existingTodo.Completed = req.Completed

	updated, err := s.Repo.Update(ctx, id, existingTodo)
	if err != nil {
		return web.TodoResponse{}, err
	}

	return helper.ToTodoResponse(updated), nil
}
func (s *TodoServiceImpl) Delete(ctx context.Context, id uuid.UUID, userId uuid.UUID) error {
	todo, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if todo.UserID != userId {
		return errors.New("Forbidden")
	}

	return s.Repo.Delete(ctx, id)
}
func (s *TodoServiceImpl) FindByID(ctx context.Context, id uuid.UUID, userId uuid.UUID) (web.TodoResponse, error) {
	todo, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.TodoResponse{}, err
	}

	if todo.UserID != userId {
		return web.TodoResponse{}, errors.New("Forbidden")
	}

	return helper.ToTodoResponse(todo), nil
}
func (s *TodoServiceImpl) FindByUserID(ctx context.Context, userId uuid.UUID) ([]web.TodoResponse, error) {
	todos, err := s.Repo.FindByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	var responses []web.TodoResponse
	for _, todo := range todos {
		responses = append(responses, helper.ToTodoResponse(todo))
	}

	return responses, nil
}
