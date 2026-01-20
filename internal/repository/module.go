package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jne100/golang-service-layout/internal/config"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRepository),
)

type Params struct {
	fx.In
	Cfg config.InventoryConfig
}

func NewRepository(p Params) (Repository, error) {
	db, err := sql.Open("sqlite3", p.Cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	r := &repository{
		cfg: p.Cfg,
		db:  db,
	}

	if err := r.initSchema(ctx); err != nil {
		return nil, err
	}

	return r, nil
}
