package templatify

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGoTemplatify(t *testing.T) {
	templateFile, err := ioutil.TempFile("", "templatify_test")
	assert.Nil(t, err)
	testConfig := config.Configuration{
		TemplateDir: filepath.Dir(templateFile.Name()),
	}
	defer os.Remove(templateFile.Name())

	templateFile.WriteString("Hello {{ .Name }}!")

	templatify := GoTemplatify{
		Config: testConfig,
	}
	data := struct {
		Name string
	}{
		Name: "World",
	}
	resultPath, err := templatify.ApplyTemplate(filepath.Base(templateFile.Name()), data)
	assert.Nil(t, err)
	defer os.Remove(resultPath)

	resultFile, err := os.Open(resultPath)
	assert.Nil(t, err)

	result, err := ioutil.ReadAll(resultFile)
	assert.Nil(t, err)

	assert.Equal(t, "Hello World!", string(result))
}
