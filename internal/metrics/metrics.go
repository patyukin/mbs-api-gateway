package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=Counter --output=../handler/mocks
type Counter interface {
	Inc()
}

var (
	IncomingTraffic Counter

	TotalRegistrations      Counter
	SuccessfulRegistrations Counter
	FailedRegistrations     Counter

	TotalLogin      Counter
	SuccessfulLogin Counter
	FailedLogin     Counter

	TotalLogReport      Counter
	SuccessfulLogReport Counter
	FailedLogReport     Counter
)

func RegisterMetrics() error {
	var err error

	IncomingTraffic = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "incoming_traffic", Help: "Incoming traffic to the application"},
	)
	err = prometheus.Register(IncomingTraffic.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register incoming traffic: %w", err)
	}

	SuccessfulRegistrations = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "successful_sign_up_v1", Help: "Successful Registrations"},
	)
	err = prometheus.Register(SuccessfulRegistrations.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register successful registrations: %w", err)
	}

	TotalRegistrations = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "total_sign_up_v1", Help: "Total registrations"},
	)
	err = prometheus.Register(TotalRegistrations.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register total registrations: %w", err)
	}

	FailedRegistrations = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "failed_sign_up_v1", Help: "Failed registrations"},
	)
	err = prometheus.Register(FailedRegistrations.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register failed registrations: %w", err)
	}

	SuccessfulLogin = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "successful_sign_in_v1", Help: "Successful login"},
	)
	err = prometheus.Register(SuccessfulLogin.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to Successful login: %w", err)
	}

	FailedLogin = prometheus.NewCounter(prometheus.CounterOpts{Name: "failed_sign_in_v1", Help: "Failed login"})
	err = prometheus.Register(FailedLogin.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to login failed login: %w", err)
	}

	TotalLogReport = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "total_v1_log_report", Help: "Total Log report"},
	)
	err = prometheus.Register(TotalLogReport.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register total log report: %w", err)
	}

	SuccessfulLogReport = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "success_v1_log_report", Help: "Success Log report"},
	)
	err = prometheus.Register(SuccessfulLogReport.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register success log report: %w", err)
	}

	FailedLogReport = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "failed_v1_log_report", Help: "Failed Log report"},
	)
	err = prometheus.Register(FailedLogReport.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register failed log report: %w", err)
	}

	return nil
}
