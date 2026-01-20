package controller

import (
	"context"

	"github.com/jne100/golang-service-layout/internal/model"
)

func (c *controller) GetItem(ctx context.Context, sku string) (model.Item, error) {
	return c.repo.ReadItem(ctx, sku)
}
