package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/abdollahpour/micro-pdf-generator/internal/generator"
	"github.com/abdollahpour/micro-pdf-generator/internal/templatify"
	"github.com/google/jsonapi"
	log "github.com/sirupsen/logrus"
)

type HttpHandler struct {
	generator  generator.Generator
	templatify templatify.Templatify
	tempDir    string
	maxSize    int
}

func NewHttpHandler(generator generator.Generator, templatify templatify.Templatify, tempDir string, maxSize int) http.Handler {
	return HttpHandler{
		generator:  generator,
		templatify: templatify,
		tempDir:    tempDir,
		maxSize:    maxSize,
	}
}

func (p HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m1 := regexp.MustCompile("\\/pdf\\/([a-z0-9]+).pdf")
	match := m1.FindStringSubmatch(strings.ToLower(req.URL.Path))
	if len(match) != 2 {
		http.Error(w, "Not found", 404)
		return
	}

	if req.Method == http.MethodPost {
		// Ignore if it's not multipart
		_ = req.ParseMultipartForm(int64(p.maxSize) << 20)
	}

	templateData, err := fetchParam(req, "template")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title:  "Failed to fetch template data",
			Detail: "You can either use Form file, field, a valid public URL or query string value to pass `template`",
			Status: "400",
			Code:   "REQ-100",
		}})
		return
	}
	if len(templateData) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title:  "Failed to fetch template data",
			Detail: "You can either use Form file, field, a valid public URL or query string value to pass `template`",
			Status: "400",
			Code:   "REQ-101",
		}})
		return
	}

	data, err := fetchParam(req, "data")
	if err != nil {
		log.WithError(err).WithField("data", data).Warn("Failed to fetch data field")
	}

	var templateFile string
	if len(data) > 0 {
		var jsonData interface{}
		err := json.NewDecoder(strings.NewReader(data)).Decode(&jsonData)
		fmt.Println(jsonData)
		if err != nil {
			log.WithError(err).WithField("data", data).Warn("'data' param is in valid JSON")
			w.WriteHeader(http.StatusBadRequest)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "'data' param is in valid JSON",
				Detail: "Faile to parse JSON data of 'data' param. Check the data format.",
				Status: "400",
				Code:   "REQ-102",
			}})
			return
		}
		templateFile, err = p.templatify.ApplyTemplate(templateData, jsonData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Failed to process the template",
				Detail: "Failed to process process template or template data",
				Status: "400",
				Code:   "REQ-103",
			}})
			return
		}
		defer os.Remove(templateFile)
	} else {
		tempFile, err := ioutil.TempFile(p.tempDir, "*.html")
		if err != nil {
			log.Fatal(err)
		}
		templateFile = tempFile.Name()
		//defer os.Remove(templateFile)

		tempFile.WriteString(templateData)
	}

	waitFor, _ := fetchParam(req, "waitFor")
	if len(waitFor) == 0 {
		waitFor = "body"
	}

	pdfFile, err := p.generator.RenderHTMLFile(templateFile, waitFor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title:  "Rendering failed",
			Detail: "Failed to render PDF from generate HTML template",
			Status: "500",
			Code:   "REQ-104",
		}})
		return
	}
	defer os.Remove(pdfFile)

	download, _ := fetchParam(req, "download")
	if strings.ToLower(download) == "true" {
		w.Header().Set("Content-Disposition", "attachment; filename="+templateData+".pdf")
	}
	http.ServeFile(w, req, pdfFile)
}

func fetchParam(r *http.Request, name string) (string, error) {
	content := param(r, name)
	if strings.HasPrefix(content, "http") {
		_, err := url.ParseRequestURI(content)
		if err == nil {
			res, err := http.Get(content)
			if err != nil {
				return "", err
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err == nil {
				return string(body), nil
			}
		}
	}

	return content, nil
}

func param(r *http.Request, name string) string {
	if r.Method == http.MethodPost {
		f, _, err := r.FormFile(name)
		if err == nil {
			defer f.Close()

			contents, err := ioutil.ReadAll(f)
			if err != nil {
				log.WithField("name", name).Error("Failed to read param file data")
				return ""
			}
			return string(contents)
		}
	}

	if r.Method == http.MethodPost {
		formValue := r.FormValue(name)
		if len(formValue) > 0 {
			return formValue
		}
	}

	return r.URL.Query().Get(name)
}
