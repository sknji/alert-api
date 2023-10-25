package service

import (
	"context"
	"github.com/sknji/alert-api/internal/persist"
)

type key int

const (
	keyStore key = 1
)

func WithStorage(ctx context.Context, s persist.Store) context.Context {
	return context.WithValue(ctx, keyStore, s)
}

func Storage(ctx context.Context) persist.Store {
	v, _ := ctx.Value(keyStore).(persist.Store)
	return v
}
