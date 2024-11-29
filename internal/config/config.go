package config

import (
	"fmt"

	configLoader "github.com/patyukin/mbs-pkg/pkg/config"
)

type Config struct {
	MinLogLevel string `validate:"oneof=debug info warn error fatal panic" yaml:"min_log_level"`
	HttpServer  struct {
		Port         int `validate:"min=1,max=65535" yaml:"port"`
		ReadTimeout  int `validate:"min=1,max=65535" yaml:"read_timeout"`
		WriteTimeout int `validate:"min=1,max=65535" yaml:"write_timeout"`
		RateLimit    struct {
			Rps   float64 `validate:"min=1,max=65535" yaml:"rps"`
			Burst int     `validate:"min=1,max=65535" yaml:"burst"`
		} `yaml:"rate_limit"`
	} `yaml:"server"`
	JwtSecret  string `validate:"required" yaml:"jwt_secret"`
	TracerHost string `validate:"required" yaml:"tracer_host"`
	GRPC       struct {
		AuthService    string `validate:"required" yaml:"auth_service"`
		LoggerService  string `validate:"required" yaml:"logger_service"`
		PaymentService string `validate:"required" yaml:"payment_service"`
		CreditService  string `validate:"required" yaml:"credit_service"`
		ReportService  string `validate:"required" yaml:"report_service"`
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
