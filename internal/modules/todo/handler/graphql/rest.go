package graphql

import (
	"context"
	"fmt"
	"graph-demo/internal/dto"
	"graph-demo/internal/modules/todo"
)

type handler struct {
	ctx     context.Context
	usecase todo.Usecase
}

func NewHandler(ctx context.Context, usecase todo.Usecase) todo.Handler {
	return &handler{
		ctx:     ctx,
		usecase: usecase,
	}
}

var todos []*dto.Todo

// CreateTodo is the resolver for the createTodo field.
func (h *handler) CreateTodo(ctx context.Context, input dto.NewTodo) (*dto.Todo, error) {
	todos = append(todos, &dto.Todo{
		ID:   fmt.Sprintf("T%d", len(todos)+1),
		Text: input.Text,
	})

	return &dto.Todo{
		ID:   fmt.Sprintf("T%d", len(todos)),
		Text: input.Text,
		Done: false,
	}, nil
}

func (h *handler) Todos(ctx context.Context) ([]*dto.Todo, error) {
	return todos, nil
}
