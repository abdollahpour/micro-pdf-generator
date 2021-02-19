package metric

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusMetric struct {
	histogram *prometheus.HistogramVec
	handler   http.Handler
}

func NewPrometheusMetric() PrometheusMetric {
	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_time",
		Help: "Time it has taken to retrieve the metrics",
	}, []string{"status", "path"})

	histogram := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "golang",
			Name:      "my_histogram",
			Help:      "This is my histogram",
		})

	prometheus.Register(histogramVec)

	prometheus.MustRegister(histogram)

	return PrometheusMetric{
		histogram: histogramVec,
		handler:   promhttp.Handler(),
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (p *PrometheusMetric) CaptureMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, req)

		p.histogram.
			WithLabelValues(fmt.Sprintf("%d", lrw.statusCode), req.URL.Path).
			Observe(time.Since(start).Seconds())
	})
	// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	// 		start := time.Now()
	// 		status := http.StatusOK

	// 		defer func() {
	// 			p.histogram.
	// 				WithLabelValues(fmt.Sprintf("%d", status), req.URL.Path).
	// 				Observe(time.Since(start).Seconds())
	// 		}()

	// 		if req.Method == http.MethodPost {
	// 			next.ServeHTTP(w, req)
	// 			return
	// 		}
	// 		status = http.StatusBadRequest

	// 		w.WriteHeader(status)
	// 	})
}

func (p *PrometheusMetric) MetricsHandler() http.Handler {
	return p.handler
}
