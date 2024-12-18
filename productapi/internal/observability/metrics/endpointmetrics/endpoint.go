package endpointmetrics

import (
	"fmt"
	"net/http"

	"github.com/odanaraujo/stocks-api/internal/observability/metrics/countermetrics"
	"github.com/odanaraujo/stocks-api/internal/observability/metrics/histogrammetrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	endpoint            string = "endpoint"
	verb                string = "verb"
	pattern             string = "pattern"
	failed              string = "failed"
	error               string = "error"
	responseCode        string = "response_code"
	isAvailabilityError string = "is_availability_error"
	isReliabilityError  string = "is_reliability_error"

	// Names
	endpointRequestCounter string = "endpoint_request_counter"
	endpointRequestLatency string = "endpoint_request_latency"
)

type Metrics struct {
	// Metric
	Latency float64

	// Labels
	Endpoint             string
	Verb                 string
	Pattern              string
	Responsecode         int
	Failed               bool
	Error                string
	HasAvailabilityError bool
	HasReliabilityError  bool
}

func Send(metrics Metrics) {
	labels := map[string]string{
		endpoint:            metrics.Endpoint,
		verb:                metrics.Verb,
		pattern:             metrics.Pattern,
		responseCode:        fmt.Sprintf("%d", metrics.Responsecode),
		failed:              fmt.Sprintf("%v", metrics.Failed),
		error:               metrics.Error,
		isAvailabilityError: fmt.Sprintf("%v", metrics.HasAvailabilityError),
		isReliabilityError:  fmt.Sprintf("%v", metrics.HasReliabilityError),
	}

	countermetrics.Increment(countermetrics.Metric{
		Name:   endpointRequestCounter,
		Labels: labels,
	})

	histogrammetrics.Observer(histogrammetrics.Metric{
		Name:  endpointRequestLatency,
		Value: float64(metrics.Latency),
		Labels: map[string]string{
			endpoint: metrics.Endpoint,
		},
	})

}

func Start() {
	fmt.Println("starting prometheus")
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		http.ListenAndServe(":2112", nil)
	}()

	fmt.Println("started prometheus")
}
