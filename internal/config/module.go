package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
)

func NewConfig() (InventoryConfig, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile(viper.GetString("CONFIG_PATH"))

	if err := viper.ReadInConfig(); err != nil {
		return InventoryConfig{}, err
	}

	var cfg InventoryConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return InventoryConfig{}, err
	}

	return cfg, nil
}
