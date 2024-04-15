package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/saufiroja/graphql-go-example/internal/delivery/graph/model"
	"github.com/saufiroja/graphql-go-example/internal/repositories"
)

type TodoService struct {
	TodoRepository repositories.ITodoRepository
}

func NewTodoService(todoRepository repositories.ITodoRepository) ITodoService {
	return &TodoService{TodoRepository: todoRepository}
}

func (s *TodoService) FindTodos() ([]*model.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todos, err := s.TodoRepository.FindTodos(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) FindTodoByTodoId(todoID string) (*model.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todo, err := s.TodoRepository.FindTodoByTodoId(todoID, ctx)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) InsertTodo(request *model.RequestInsertTodo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := uuid.New().String()
	request.TodoID = &id

	err := s.TodoRepository.InsertTodo(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) UpdateTodoByTodoId(todoID string, request *model.RequestUpdateTodo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.TodoRepository.FindTodoByTodoId(todoID, ctx)
	if err != nil {
		return err
	}

	err = s.TodoRepository.UpdateTodoByTodoId(ctx, todoID, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) DeleteTodoByTodoId(todoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.TodoRepository.FindTodoByTodoId(todoID, ctx)
	if err != nil {
		return err
	}

	err = s.TodoRepository.DeleteTodoByTodoId(ctx, todoID)
	if err != nil {
		return err
	}

	return nil
}
