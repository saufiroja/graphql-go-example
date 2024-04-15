package repositories

import (
	"context"

	"github.com/saufiroja/graphql-go-example/internal/delivery/graph/model"
)

type ITodoRepository interface {
	FindTodos(ctx context.Context) ([]*model.Todo, error)
	FindTodoByTodoId(todoID string, ctx context.Context) (*model.Todo, error)
	InsertTodo(ctx context.Context, request *model.RequestInsertTodo) error
	UpdateTodoByTodoId(ctx context.Context, todoID string, request *model.RequestUpdateTodo) error
	DeleteTodoByTodoId(ctx context.Context, todoID string) error
}
