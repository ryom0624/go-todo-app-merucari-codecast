package db

import (
	"context"
	"go-todoapp/internal/todo"
)

type DB interface {
	PutTODO(ctx context.Context, t *todo.TODO) error
}
