package cron

import (
	"context"
	"time"

	"github.com/jne100/golang-service-layout/internal/config"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Invoke(RegisterJobs),
)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Cfg       config.InventoryConfig
}

func RegisterJobs(p Params) error {
	loc, err := time.LoadLocation(p.Cfg.Cron.Timezone)
	if err != nil {
		return err
	}

	c := cron.New(cron.WithLocation(loc))
	c.AddFunc(p.Cfg.Cron.PrintStats, func() {
		zap.L().Info("print stats here...")
	})

	c.Start()

	p.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			c.Stop()
			return nil
		},
	})

	return nil
}
