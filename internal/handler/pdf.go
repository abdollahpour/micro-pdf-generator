package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/abdollahpour/micro-pdf-generator/internal/pdf"
	"github.com/abdollahpour/micro-pdf-generator/internal/templatify"
)

type PdfHandler struct {
	PdfGenerator pdf.PdfGenerator
	Templatify   templatify.Templatify
}

func (p PdfHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	m1 := regexp.MustCompile(`[^a-z0-9]`)
	templateName := m1.ReplaceAllString(strings.ToLower(req.URL.Path), "")

	var templateData interface{}
	err := json.NewDecoder(req.Body).Decode(&templateData)
	if err != nil {
		log.Println(fmt.Sprintf("Error ro parse the JSON %v", req.Body))
		http.Error(res, "Illegal json", 400)
		return
	}

	templateFile, err := p.Templatify.ApplyTemplate(templateName, templateData)
	if err != nil {
		switch e := err.(type) {
		case *templatify.TemplateError:
			if e.NotFound {
				http.Error(res, e.Error(), 404)
			} else if !e.Processed {
				http.Error(res, e.Error(), 400)
			} else {
				http.Error(res, e.Error(), 500)
			}
		}
		return
	}
	defer os.Remove(templateFile)

	pdfFile, err := p.PdfGenerator.RenderUrlToPdf(templateFile, "#main")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	defer os.Remove(pdfFile)

	http.ServeFile(res, req, pdfFile)
}
