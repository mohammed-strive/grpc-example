package graph

import (
	"aeon-grpc/graph/model"
	"aeon-grpc/interfaces"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Store interfaces.StoreClient[model.Book]
}
