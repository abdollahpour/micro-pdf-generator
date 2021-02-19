package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/abdollahpour/micro-pdf-generator/internal/handler"
	"github.com/abdollahpour/micro-pdf-generator/internal/metric"
	"github.com/abdollahpour/micro-pdf-generator/internal/pdf"
	"github.com/abdollahpour/micro-pdf-generator/internal/templatify"
)

func main() {
	config := config.EnvConfig
	pdfGenerator := pdf.ChromedpGenerator{Config: config}
	template := templatify.GoTemplatify{Config: config}
	pdfHandler := handler.PdfHandler{PdfGenerator: pdfGenerator, Templatify: template}
	metric := metric.NewPrometheusMetric()

	mux := http.NewServeMux()
	mux.Handle("/metrics", metric.MetricsHandler())
	mux.Handle("/pdf/", metric.CaptureMetrics(pdfHandler))
	log.Println(fmt.Sprintf("Listen on port %d", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), mux))
}
