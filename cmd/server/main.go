package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/abdollahpour/micro-pdf-generator/internal/generator"
	"github.com/abdollahpour/micro-pdf-generator/internal/handler"
	"github.com/abdollahpour/micro-pdf-generator/internal/metric"
	"github.com/abdollahpour/micro-pdf-generator/internal/templatify"
	log "github.com/sirupsen/logrus"
)

var Version = "development"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	version := flag.Bool("version", false, "print version version")
	debug := flag.Bool("debug", false, "active debug manager")

	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	if *debug {
		log.SetLevel(log.TraceLevel)
	}

	conf := config.NewEnvConfiguration()

	if _, err := os.Stat(conf.TempDir); os.IsNotExist(err) {
		err = os.Mkdir(conf.TempDir, 0744)
		if err != nil {
			log.Fatal("Failed to create temp dir: " + conf.TempDir)
		}
	}

	pdfGenerator := generator.NewChromedpGenerator(conf.Timeout, conf.TempDir)
	template := templatify.NewGoTemplatify(conf.TempDir)
	httpHandler := handler.NewHttpHandler(pdfGenerator, template, conf.TempDir, conf.MaxSize)
	metric := metric.NewPrometheusMetric()

	mux := http.NewServeMux()
	mux.Handle("/metrics", metric.MetricsHandler())
	mux.Handle("/pdf/", metric.CaptureMetrics(httpHandler))
	log.Println(fmt.Sprintf("Listen on port %d", conf.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), mux))
}
