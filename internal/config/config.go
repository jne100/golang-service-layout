package config

type InventoryConfig struct {
	Inbound Inbound `toml:"inbound"`
	Logging Logging `toml:"logging"`
	Db      Db      `toml:"db"`
	Cron    Cron    `toml:"cron"`
}

type Inbound struct {
	GrpcPort int `toml:"grpc_port"`
}

type Logging struct {
	Level string `toml:"level"`
}

type Db struct {
	Driver string `toml:"driver"`
	Dsn    string `toml:"dsn"`
}

type Cron struct {
	Timezone   string `toml:"timezone"`
	PrintStats string `toml:"print_stats"`
}
