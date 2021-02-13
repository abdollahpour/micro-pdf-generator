package templatify

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type TemplateError struct {
	NotFound  bool
	Processed bool
	msg       string
}

func (e *TemplateError) Error() string { return e.msg }

func ApplyTemplate(name string, data interface{}) (string, error) {
	templateFile := name + ".html"
	_, err := os.Stat(templateFile)
	if os.IsNotExist(err) {
		return "", &TemplateError{
			NotFound: true,
			msg:      fmt.Sprintf("Template %v not found (%v)", name, templateFile),
		}
	}

	teml, err := template.New(templateFile).ParseFiles("./" + templateFile)
	if err != nil {
		log.Print(err)
		return "", &TemplateError{
			Processed: false,
			msg:       fmt.Sprintf("Error to parse %v", name),
		}
	}

	templateBuf := new(bytes.Buffer)
	err = teml.Execute(templateBuf, data)
	if err != nil {
		log.Print(err)
		return "", &TemplateError{msg: fmt.Sprintf("Error to process the template %v", name)}
	}

	tempFile := "temp_" + strconv.FormatInt(time.Now().UnixNano(), 16) + ".html"

	err = ioutil.WriteFile(tempFile, []byte(templateBuf.String()), 0644)
	if err != nil {
		log.Print(err)
		return "", &TemplateError{msg: fmt.Sprintf("Eror to write template file %v", name)}
	}

	return tempFile, nil
}
