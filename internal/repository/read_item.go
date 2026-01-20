package repository

import (
	"context"

	"github.com/jne100/golang-service-layout/internal/model"
)

func (r *repository) ReadItem(ctx context.Context, sku string) (model.Item, error) {
	var item model.Item

	err := r.db.QueryRowContext(
		ctx,
		`SELECT sku, name, quantity
		 FROM items
		 WHERE sku = ?`,
		sku,
	).Scan(&item.Sku, &item.Name, &item.Quantity)

	return item, err
}
