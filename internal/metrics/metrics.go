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
	IncomingTraffic         Counter
	SuccessfulRegistrations Counter
	TotalRegistrations      Counter
	FailedRegistrations     Counter
)

func RegisterMetrics() error {
	var err error

	IncomingTraffic = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "incoming_traffic",
		Help: "Incoming traffic to the application",
	})

	err = prometheus.Register(IncomingTraffic.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register incoming traffic: %w", err)
	}

	SuccessfulRegistrations = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "successful_sign_up_v1",
		Help: "Successful Registrations",
	})

	err = prometheus.Register(SuccessfulRegistrations.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register successful registrations: %w", err)
	}

	TotalRegistrations = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_sign_up_v1",
		Help: "Total registrations",
	})

	err = prometheus.Register(TotalRegistrations.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register total registrations: %w", err)
	}

	FailedRegistrations = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "failed_sign_up_v1",
		Help: "Failed registrations",
	})

	err = prometheus.Register(FailedRegistrations.(prometheus.Collector))
	if err != nil {
		return fmt.Errorf("failed to register failed registrations: %w", err)
	}

	return nil
}
