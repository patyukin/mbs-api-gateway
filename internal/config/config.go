package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	MinLogLevel string `yaml:"min_log_level" validate:"oneof=debug info warn error fatal panic"`
	HttpServer  struct {
		Port         int `yaml:"port" validate:"min=1,max=65535"`
		ReadTimeout  int `yaml:"read_timeout" validate:"min=1,max=65535"`
		WriteTimeout int `yaml:"write_timeout" validate:"min=1,max=65535"`
		RateLimit    int `yaml:"rate_limit" validate:"min=1,max=65535"`
	} `yaml:"server"`
	JwtSecret  string `yaml:"jwt_secret" validate:"required"`
	TracerHost string `yaml:"tracer_host" validate:"required"`
	GRPC       struct {
		AuthServicePort int `yaml:"auth_service_port" validate:"min=1,max=65535"`
	} `yaml:"grpc"`
}

func LoadConfig() (*Config, error) {
	yamlConfigFilePath := os.Getenv("YAML_CONFIG_FILE_PATH")
	if yamlConfigFilePath == "" {
		return nil, fmt.Errorf("yaml config file path is not set")
	}

	f, err := os.Open(yamlConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %w", err)
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Printf("unable to close config file: %v", err)
		}
	}(f)

	var config Config
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config file: %w", err)
	}

	validate := validator.New()
	if err = validate.Struct(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}
