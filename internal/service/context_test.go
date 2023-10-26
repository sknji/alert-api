package service

import (
	"context"
	"github.com/sknji/alert-api/internal/persist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()

	expected := persist.NewDummyStore()
	ctx = WithStorage(ctx, expected)
	found := Storage(ctx)
	assert.Equal(t, expected, found)
}

func TestAlertService(t *testing.T) {
	ctx := context.Background()

	expected := &AlertServiceImpl{}
	ctx = WithAlertService(ctx, expected)
	found := AlertService(ctx)
	assert.Equal(t, expected, found)
}
