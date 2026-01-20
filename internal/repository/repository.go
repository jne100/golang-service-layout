package repository

import (
	"context"
	"database/sql"

	"github.com/jne100/golang-service-layout/internal/config"
	"github.com/jne100/golang-service-layout/internal/model"
)

type Repository interface {
	InsertItem(ctx context.Context, item model.Item) error
	ReadItem(ctx context.Context, sku string) (model.Item, error)
}

type repository struct {
	cfg config.InventoryConfig
	db  *sql.DB
}
