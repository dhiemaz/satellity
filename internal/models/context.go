package models

import (
	"context"
	"satellity/internal/durable"
)

// Context application
type Context struct {
	context  context.Context
	database *durable.Database
}

// WrapContext application
func WrapContext(ctx context.Context, db *durable.Database) *Context {
	return &Context{context: ctx, database: db}
}
