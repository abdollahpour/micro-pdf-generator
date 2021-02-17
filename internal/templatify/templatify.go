package templatify

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
)

type Templatify interface {
	ApplyTemplate(name string, data interface{}) (string, error)
}

type TemplateError struct {
	NotFound  bool
	Processed bool
	msg       string
}

func (e *TemplateError) Error() string { return e.msg }

// GoTemplatify is using go built-in template engine
type GoTemplatify struct {
	Config config.Configuration
}

func (g GoTemplatify) ApplyTemplate(templateName string, data interface{}) (string, error) {
	templateFile := filepath.FromSlash(g.Config.TemplateDir + "/" + templateName)
	_, err := os.Stat(templateFile)
	if os.IsNotExist(err) {
		return "", &TemplateError{
			NotFound: true,
			msg:      fmt.Sprintf("Template %v not found (%v)", templateName, templateFile),
		}
	}

	teml, err := template.New(templateName).ParseFiles(templateFile)
	if err != nil {
		log.Print(err)
		return "", &TemplateError{
			Processed: false,
			msg:       fmt.Sprintf("Error to parse %v", templateName),
		}
	}

	templateBuf := new(bytes.Buffer)
	err = teml.Execute(templateBuf, data)
	if err != nil {
		log.Print(err)
		return "", &TemplateError{msg: fmt.Sprintf("Error to process the template %v", templateName)}
	}

	tempFile := "temp_" + strconv.FormatInt(time.Now().UnixNano(), 16) + ".html"

	err = ioutil.WriteFile(tempFile, []byte(templateBuf.String()), 0644)
	if err != nil {
		log.Print(err)
		return "", &TemplateError{msg: fmt.Sprintf("Eror to write template file %v", templateName)}
	}

	return tempFile, nil
}
