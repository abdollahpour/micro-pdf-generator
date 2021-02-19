package metric

import (
	"net/http"
)

type Metric interface {
	CaptureMetrics(next http.Handler) http.Handler
	MetricsHandler() http.Handler
}
