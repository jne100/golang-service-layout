package controller

import (
	"context"

	"github.com/jne100/golang-service-layout/internal/model"
	"github.com/jne100/golang-service-layout/internal/repository"
)

type Controller interface {
	CreateItem(ctx context.Context, item model.Item) error
	GetItem(ctx context.Context, sku string) (model.Item, error)
}

type controller struct {
	repo repository.Repository
}
