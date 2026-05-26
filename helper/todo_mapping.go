package helper

import (
	"golang-blog-api/domain"
	"golang-blog-api/model/web"
)

func ToTodoResponse(todo domain.Todo) web.TodoResponse {
	return web.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

// func ToTodoResponses(todo domain.Todo) web.TodoResponse {
// 	return web.TodoResponse{
// 		ID:          todo.ID,
// 		Title:       todo.Title,
// 		Description: todo.Description,
// 		Completed:   todo.Completed,
// 		CreatedAt:   todo.CreatedAt,
// 		UpdatedAt:   todo.UpdatedAt,
// 	}
// }
