package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define a counter metric
	requestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed.",
		},
		[]string{"method", "status"},
	)

	// Define a gauge metric
	responseDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_response_duration_seconds",
			Help: "Duration of HTTP requests in seconds.",
		},
		[]string{"method"},
	)

	// Define a histogram metric
	durationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_duration_seconds",
			Help:    "Histogram of response durations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(requestsCounter)
	prometheus.MustRegister(responseDuration)
	prometheus.MustRegister(durationHistogram)
}

func main() {

	// Expose the metrics at "/metrics"
	http.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
