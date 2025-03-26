package config

type Config struct {
	HTTP struct {
		Port string `env:"HTTP_PORT" env-default:"8080"`
	}
	Postgres struct {
		DSN string `env:"PG_DSN,required"`
	}
	Redis struct {
		Addr     string `env:"REDIS_ADDR" env-default:"localhost:6379"`
		Password string `env:"REDIS_PASSWORD" env-default:""`
		DB       int    `env:"REDIS_DB" env-default:"0"`
	}
}
