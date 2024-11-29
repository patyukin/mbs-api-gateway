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

func RegisterLoggerMetrics() error {
	var err error

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
