package controller

import (
	"context"

	"github.com/jne100/golang-service-layout/internal/model"
)

func (c *controller) CreateItem(ctx context.Context, item model.Item) error {
	return c.repo.InsertItem(ctx, item)
}
