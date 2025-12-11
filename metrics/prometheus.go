package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

var (
	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Request duration in seconds",
			Buckets: prometheus.DefBuckets, // Use default buckets
		},
		[]string{"method", "endpoint"},
	)
)

// Init registers the metrics with Prometheus.
func Init() {
	prometheus.MustRegister(TotalRequests)
	prometheus.MustRegister(RequestDuration)
}

// MetricsMiddleware records the request metrics.
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Pass to the next handler
		next.ServeHTTP(w, r)

		// Record metrics after handler execution
		duration := time.Since(start).Seconds()
		TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()
		RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}
