package services

import "github.com/saufiroja/graphql-go-example/internal/delivery/graph/model"

type ITodoService interface {
	FindTodos() ([]*model.Todo, error)
	FindTodoByTodoId(todoID string) (*model.Todo, error)
	InsertTodo(request *model.RequestInsertTodo) error
	UpdateTodoByTodoId(todoID string, request *model.RequestUpdateTodo) error
	DeleteTodoByTodoId(todoID string) error
}
