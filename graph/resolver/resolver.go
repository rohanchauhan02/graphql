package graph

import (
	"graph-demo/internal/modules/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoHandler todo.Handler
}
