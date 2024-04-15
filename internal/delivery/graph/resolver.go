package graph

import "github.com/saufiroja/graphql-go-example/internal/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoService services.ITodoService
}
