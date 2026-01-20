package repository

import (
	"context"

	"github.com/jne100/golang-service-layout/internal/model"
)

func (r *repository) InsertItem(ctx context.Context, item model.Item) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO items (sku, name, quantity)
		 VALUES (?, ?, ?)`,
		item.Sku,
		item.Name,
		item.Quantity,
	)
	return err
}
