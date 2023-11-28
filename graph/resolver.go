package graph

import "go.uber.org/zap"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log *zap.Logger
}

func NewResolver(
	log *zap.Logger,
) *Resolver {
	return &Resolver{
		log: log,
	}
}
