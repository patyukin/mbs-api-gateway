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

	return nil
}
