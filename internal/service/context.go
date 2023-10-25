package service

import (
	"context"
	"github.com/sknji/alert-api/internal/persist"
)

type key int

const (
	keyStore key = iota
	keyAlertService
)

func WithStorage(ctx context.Context, s persist.Store) context.Context {
	return context.WithValue(ctx, keyStore, s)
}

func Storage(ctx context.Context) persist.Store {
	v, _ := ctx.Value(keyStore).(persist.Store)
	return v
}

func WithAlertService(ctx context.Context, a AlertServiceI) context.Context {
	return context.WithValue(ctx, keyAlertService, a)
}

func AlertService(ctx context.Context) AlertServiceI {
	v, _ := ctx.Value(keyAlertService).(AlertServiceI)
	return v
}
