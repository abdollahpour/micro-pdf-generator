package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/abdollahpour/micro-pdf-generator/internal/generator"
	"github.com/abdollahpour/micro-pdf-generator/internal/handler"
	"github.com/abdollahpour/micro-pdf-generator/internal/metric"
	"github.com/abdollahpour/micro-pdf-generator/internal/templatify"
)

func main() {
	config := config.EnvConfig
	pdfGenerator := generator.NewChromedpGenerator(config.Timeout, config.TempDir)
	template := templatify.NewGoTemplatify(config.TempDir)
	httpHandler := handler.NewHttpHandler(pdfGenerator, template, config.TempDir, config.MaxSize)
	metric := metric.NewPrometheusMetric()

	mux := http.NewServeMux()
	mux.Handle("/metrics", metric.MetricsHandler())
	mux.Handle("/pdf/", metric.CaptureMetrics(httpHandler))
	log.Println(fmt.Sprintf("Listen on port %d", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), mux))
}
