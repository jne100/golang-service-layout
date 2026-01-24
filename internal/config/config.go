package config

type InventoryConfig struct {
	Inbound Inbound `mapstructure:"inbound"`
	Logging Logging `mapstructure:"logging"`
	Db      Db      `mapstructure:"db"`
	Cron    Cron    `mapstructure:"cron"`
}

type Inbound struct {
	GrpcPort int `mapstructure:"grpc_port"`
}

type Logging struct {
	Level string `mapstructure:"level"`
}

type Db struct {
	Driver string `mapstructure:"driver"`
	Dsn    string `mapstructure:"dsn"`
}

type Cron struct {
	Timezone   string `mapstructure:"timezone"`
	PrintStats string `mapstructure:"print_stats"`
}
