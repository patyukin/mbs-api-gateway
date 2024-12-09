package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// Declare counters as Prometheus-compatible types.
var (
	TotalRegistrations      prometheus.Counter
	SuccessfulRegistrations prometheus.Counter
	FailedRegistrations     prometheus.Counter

	TotalLogin      prometheus.Counter
	SuccessfulLogin prometheus.Counter
	FailedLogin     prometheus.Counter
)

// RegisterAuthMetrics initializes and registers Prometheus metrics for authentication.
func RegisterAuthMetrics() error {
	var err error

	// Initialize and register the SuccessfulRegistrations counter
	SuccessfulRegistrations = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "successful_sign_up_v1",
			Help: "Successful Registrations",
		},
	)
	err = prometheus.Register(SuccessfulRegistrations)
	if err != nil {
		return fmt.Errorf("failed to register successful registrations: %w", err)
	}

	// Initialize and register the TotalRegistrations counter
	TotalRegistrations = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_sign_up_v1",
			Help: "Total registrations",
		},
	)
	err = prometheus.Register(TotalRegistrations)
	if err != nil {
		return fmt.Errorf("failed to register total registrations: %w", err)
	}

	// Initialize and register the FailedRegistrations counter
	FailedRegistrations = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "failed_sign_up_v1",
			Help: "Failed registrations",
		},
	)
	err = prometheus.Register(FailedRegistrations)
	if err != nil {
		return fmt.Errorf("failed to register failed registrations: %w", err)
	}

	// Initialize and register the TotalLogin counter
	TotalLogin = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_sign_in_v1",
			Help: "Total logins",
		},
	)
	err = prometheus.Register(TotalLogin)
	if err != nil {
		return fmt.Errorf("failed to register total logins: %w", err)
	}

	// Initialize and register the SuccessfulLogin counter
	SuccessfulLogin = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "successful_sign_in_v1",
			Help: "Successful login",
		},
	)
	err = prometheus.Register(SuccessfulLogin)
	if err != nil {
		return fmt.Errorf("failed to register successful login: %w", err)
	}

	// Initialize and register the FailedLogin counter
	FailedLogin = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "failed_sign_in_v1",
			Help: "Failed login",
		},
	)
	err = prometheus.Register(FailedLogin)
	if err != nil {
		return fmt.Errorf("failed to register failed login: %w", err)
	}

	return nil
}
