package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/saufiroja/graphql-go-example/internal/delivery/graph/model"
)

// InsertTodo is the resolver for the insertTodo field.
func (r *mutationResolver) InsertTodo(ctx context.Context, input model.RequestInsertTodo) (*model.Todo, error) {
	err := r.TodoService.InsertTodo(&input)
	if err != nil {
		return nil, err
	}

	res := &model.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   *input.Completed,
	}

	return res, nil
}

// UpdateTodoByTodoID is the resolver for the updateTodoByTodoId field.
func (r *mutationResolver) UpdateTodoByTodoID(ctx context.Context, todoID string, input model.RequestUpdateTodo) (*model.Todo, error) {
	err := r.TodoService.UpdateTodoByTodoId(todoID, &input)
	if err != nil {
		return nil, err
	}

	res := &model.Todo{
		TodoID:      todoID,
		Title:       *input.Title,
		Description: input.Description,
		Completed:   *input.Completed,
	}

	return res, nil
}

// DeleteTodoByTodoID is the resolver for the deleteTodoByTodoId field.
func (r *mutationResolver) DeleteTodoByTodoID(ctx context.Context, todoID string) (*model.Todo, error) {
	err := r.TodoService.DeleteTodoByTodoId(todoID)
	if err != nil {
		return nil, err
	}

	res := &model.Todo{
		TodoID: todoID,
	}

	return res, nil
}

// FindTodos is the resolver for the findTodos field.
func (r *queryResolver) FindTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.TodoService.FindTodos()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// FindTodoByTodoID is the resolver for the findTodoByTodoId field.
func (r *queryResolver) FindTodoByTodoID(ctx context.Context, todoID string) (*model.Todo, error) {
	todo, err := r.TodoService.FindTodoByTodoId(todoID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.RequestInsertTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.RequestUpdateTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: UpdateTodo - updateTodo"))
}
func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}
func (r *queryResolver) FindTodoByID(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: FindTodoByID - findTodoById"))
}
