package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/abdollahpour/micro-pdf-generator/internal/handler"
	"github.com/abdollahpour/micro-pdf-generator/internal/pdf"
)

func main() {
	config := config.EnvConfig
	pdfGenerator := pdf.ChromedpPdfGenerator{Config: config}
	pdfHandler := handler.PdfHandler{PdfGenerator: pdfGenerator}

	http.Handle("/", pdfHandler)
	log.Println(fmt.Sprintf("Listen on port %d", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), nil))
}
