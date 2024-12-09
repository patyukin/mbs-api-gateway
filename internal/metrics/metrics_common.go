package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	ioprometheusclient "github.com/prometheus/client_model/go"
)

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=Counter --output=../handler/mocks
type Counter interface {
	Inc()
	Desc() *prometheus.Desc
	Write(metric *ioprometheusclient.Metric) error
	Describe(descs chan<- *prometheus.Desc)
	Collect(metrics chan<- prometheus.Metric)
	Add(f float64)
}

var IncomingTraffic prometheus.Counter

func RegisterMetrics() error {
	var err error

	// Initialize and register the IncomingTraffic metric
	IncomingTraffic = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "incoming_traffic",
			Help: "Incoming traffic to the application",
		},
	)
	err = prometheus.Register(IncomingTraffic)
	if err != nil {
		return fmt.Errorf("failed to register incoming traffic: %w", err)
	}

	// Register other metric groups
	err = RegisterAuthMetrics()
	if err != nil {
		return fmt.Errorf("failed to register auth metrics: %w", err)
	}

	err = RegisterLoggerMetrics()
	if err != nil {
		return fmt.Errorf("failed to register logger metrics: %w", err)
	}

	return nil
}
