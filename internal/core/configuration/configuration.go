package configuration

import "github.com/Netflix/go-env"

type DBConfig struct {
	Host string `env:"DB_HOST"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Name string `env:"DB_NAME"`
}

type ServerConfig struct {
	HttpSocket string `env:"HTTP_SOCKET"`
	GrpcSocket string `env:"GRPC_SOCKET"`
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func Load() (*Config, error) {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
