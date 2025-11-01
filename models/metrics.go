package models

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	Name         string
	PromGauge    prometheus.Gauge
	PromGaugeVec *prometheus.GaugeVec
}
