package repositories

import (
	"context"

	"github.com/saufiroja/graphql-go-example/internal/delivery/graph/model"
	"github.com/saufiroja/graphql-go-example/pkg/database"
)

type TodoRepository struct {
	db database.IPostgres
}

func NewTodoRepository(db database.IPostgres) ITodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) FindTodos(ctx context.Context) ([]*model.Todo, error) {
	db := r.db.Db()
	query := `SELECT todo_id, title, description, completed, created_at, updated_at FROM todos`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}

func (r *TodoRepository) FindTodoByTodoId(todoID string, ctx context.Context) (*model.Todo, error) {
	db := r.db.Db()

	query := `SELECT todo_id, title, description, completed, created_at, updated_at FROM todos WHERE todo_id = $1`
	row := db.QueryRowContext(ctx, query, todoID)

	var todo model.Todo
	err := row.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) InsertTodo(ctx context.Context, request *model.RequestInsertTodo) error {
	db := r.db.Db()

	query := `INSERT INTO todos (todo_id, title, description, completed) VALUES ($1, $2, $3, $4);`
	_, err := db.ExecContext(
		ctx,
		query,
		request.TodoID,
		request.Title,
		request.Description,
		request.Completed,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) UpdateTodoByTodoId(ctx context.Context, todoID string, request *model.RequestUpdateTodo) error {
	db := r.db.Db()

	query := `UPDATE todos SET title = $1, description = $2, completed = $3 WHERE todo_id = $4;`

	_, err := db.ExecContext(
		ctx,
		query,
		request.Title,
		request.Description,
		request.Completed,
		todoID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) DeleteTodoByTodoId(ctx context.Context, todoID string) error {
	db := r.db.Db()

	query := `DELETE FROM todos WHERE todo_id = $1;`
	_, err := db.ExecContext(ctx, query, todoID)
	if err != nil {
		return err
	}

	return nil
}
