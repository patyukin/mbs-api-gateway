package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalLogReport      Counter
	SuccessfulLogReport Counter
	FailedLogReport     Counter
)

// RegisterLoggerMetrics initializes and registers logger-related metrics.
func RegisterLoggerMetrics() error {
	var err error

	// Initialize and register the TotalLogReport metric
	TotalLogReport = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_v1_log_report",
			Help: "Total Log report",
		},
	)
	err = prometheus.Register(TotalLogReport)
	if err != nil {
		return fmt.Errorf("failed to register total log report: %w", err)
	}

	// Initialize and register the SuccessfulLogReport metric
	SuccessfulLogReport = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "success_v1_log_report",
			Help: "Success Log report",
		},
	)
	err = prometheus.Register(SuccessfulLogReport)
	if err != nil {
		return fmt.Errorf("failed to register success log report: %w", err)
	}

	// Initialize and register the FailedLogReport metric
	FailedLogReport = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "failed_v1_log_report",
			Help: "Failed Log report",
		},
	)
	err = prometheus.Register(FailedLogReport)
	if err != nil {
		return fmt.Errorf("failed to register failed log report: %w", err)
	}

	return nil
}
