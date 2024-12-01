package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalRegistrations      Counter
	SuccessfulRegistrations Counter
	FailedRegistrations     Counter

	TotalLogin      Counter
	SuccessfulLogin Counter
	FailedLogin     Counter
)

func RegisterAuthMetrics() error {
	var err error

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

	TotalLogin = prometheus.NewCounter(
		prometheus.CounterOpts{Name: "total_sign_in_v1", Help: "Total logins"},
	)
	err = prometheus.Register(TotalLogin.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register total logins: %w", err)
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

	return nil
}
