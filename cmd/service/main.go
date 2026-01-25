package main

import (
	"context"

	"github.com/jne100/golang-service-layout/internal/config"
	"github.com/jne100/golang-service-layout/internal/controller"
	"github.com/jne100/golang-service-layout/internal/cron"
	"github.com/jne100/golang-service-layout/internal/handler"
	"github.com/jne100/golang-service-layout/internal/repository"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var appOpts = []fx.Option{
	handler.Module,
	controller.Module,
	repository.Module,
	config.Module,
	cron.Module,
	fx.Invoke(setVerbosity),
	fx.Invoke(sayHello),
	fx.Invoke(handler.RunHandlerAsync),
}

func main() {
	fx.New(appOpts...).Run()
}

func setVerbosity(cfg config.InventoryConfig) error {
	zapCfg := zap.NewProductionConfig()
	if err := zapCfg.Level.UnmarshalText([]byte(cfg.Logging.Level)); err != nil {
		zapCfg.Level.SetLevel(zap.InfoLevel)
	}

	logger, err := zapCfg.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)
	return nil
}

func sayHello(lc fx.Lifecycle) {
	zap.L().Info("inventory service started")
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			zap.L().Info("inventory service stopped")
			return nil
		},
	})
}
