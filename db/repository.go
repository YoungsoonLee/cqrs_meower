package db

import "context"

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
}
