package histogrammetrics

import (
	"github.com/odanaraujo/stocks-api/internal/observability/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	Name   string
	Help   string
	Value  float64
	Labels map[string]string
}

var createMetrics = make(map[string]*prometheus.HistogramVec)

func Observer(metric Metric) {
	go func() {
		labelsKeys := metrics.GetLabelsKey(metric.Labels)

		opts := prometheus.HistogramOpts{
			Name:    metric.Name,
			Help:    metric.Help,
			Buckets: GetDefaultBuckets(),
		}

		if createMetrics[metric.Name] == nil {
			histogram := promauto.NewHistogramVec(opts, labelsKeys)
			createMetrics[metric.Name] = histogram
		}

		histogram := createMetrics[metric.Name]
		histogram.With(metric.Labels).Observe(metric.Value)
	}()
}

func GetDefaultBuckets() []float64 {
	return prometheus.LinearBuckets(0.05, 0.050, 20)
}
