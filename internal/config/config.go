package config

import (
	"fmt"
	configLoader "github.com/patyukin/mbs-pkg/pkg/config"
)

type Config struct {
	MinLogLevel string `yaml:"min_log_level" validate:"oneof=debug info warn error fatal panic"`
	HttpServer  struct {
		Port         int `yaml:"port" validate:"min=1,max=65535"`
		ReadTimeout  int `yaml:"read_timeout" validate:"min=1,max=65535"`
		WriteTimeout int `yaml:"write_timeout" validate:"min=1,max=65535"`
		RateLimit    struct {
			Rps   float64 `yaml:"rps" validate:"min=1,max=65535"`
			Burst int     `yaml:"burst" validate:"min=1,max=65535"`
		} `yaml:"rate_limit"`
	} `yaml:"server"`
	JwtSecret  string `yaml:"jwt_secret" validate:"required"`
	TracerHost string `yaml:"tracer_host" validate:"required"`
	GRPC       struct {
		AuthServiceHost   string `yaml:"auth_service_host" validate:"required"`
		AuthServicePort   int    `yaml:"auth_service_port" validate:"min=1,max=65535"`
		LoggerServiceHost string `yaml:"logger_service_host" validate:"required"`
		LoggerServicePort int    `yaml:"logger_service_port" validate:"min=1,max=65535"`
	} `yaml:"grpc"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := configLoader.LoadConfig(&config)
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	return &config, nil
}
