package repository

import (
	"context"
)

func (r *repository) initSchema(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS items (
			sku TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			quantity INTEGER NOT NULL
		);
	`)
	return err
}
