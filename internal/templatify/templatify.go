package templatify

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Templatify interface {
	ApplyTemplate(templateData string, data interface{}) (string, error)
}

type TemplateError struct {
	Processed bool
	msg       string
}

// GoTemplatify is using go built-in template engine
type GoTemplatify struct {
	tempDir string
}

func NewGoTemplatify(tempDir string) *GoTemplatify {
	return &GoTemplatify{tempDir: tempDir}
}

func (g GoTemplatify) ApplyTemplate(templateData string, data interface{}) (string, error) {
	teml, err := template.New("template").Parse(templateData)
	if err != nil {
		log.WithError(err).WithField("data", templateData).Printf("Failed to parse the template")
		return "", errors.Wrap(err, "Failed to parse the template")
	}

	templateBuf := new(bytes.Buffer)
	err = teml.Execute(templateBuf, data)
	if err != nil {
		log.WithError(err).WithField("data", templateData).Printf("Failed to process the template")
		return "", errors.Wrap(err, "Failed to process the template")
	}

	tempFile, err := ioutil.TempFile(g.tempDir, "*.html")
	if err != nil {
		log.Fatalln("Failed to create temp file")
	}
	_, err = tempFile.Write([]byte(templateBuf.String()))
	if err != nil {
		log.Fatalln("Failed to write to temp file")
	}

	return tempFile.Name(), nil
}
