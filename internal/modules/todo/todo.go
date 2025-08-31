package todo

import (
	"context"
	"graph-demo/internal/dto"
)

type Handler interface {
	CreateTodo(ctx context.Context, input dto.NewTodo) (*dto.Todo, error)
	Todos(ctx context.Context) ([]*dto.Todo, error)
}

type Usecase interface {
}

type Repository interface {
}
