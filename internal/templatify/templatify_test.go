package templatify

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoTemplatify(t *testing.T) {
	templatify := NewGoTemplatify(os.TempDir())
	data := struct {
		Name string
	}{
		Name: "World",
	}
	resultPath, err := templatify.ApplyTemplate("Hello {{ .Name }}!", data)
	assert.Nil(t, err)
	defer os.Remove(resultPath)

	resultFile, err := os.Open(resultPath)
	assert.Nil(t, err)

	result, err := ioutil.ReadAll(resultFile)
	assert.Nil(t, err)

	assert.Equal(t, "Hello World!", string(result))
}
