package pdf

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestChromedpGenerator(t *testing.T) {
	tempFile, err := os.Create(path.Join(config.EnvConfig.TempDir, "TestChromedpGenerator.html"))
	assert.Nil(t, err)
	defer os.Remove(tempFile.Name())

	htmlFile, err := os.Create(tempFile.Name())
	assert.Nil(t, err)

	err = ioutil.WriteFile(htmlFile.Name(), []byte(`<!DOCTYPE html>
		<html lang="en">
		
		<head>
		</head>
		
		<body>
		  <div id="main">
			This is a test
		  </div>
		</body>
		
		</html>
	`), 0777)
	assert.Nil(t, err)

	chromedpGenerator := ChromedpGenerator{
		Config: config.EnvConfig,
	}
	pdfFile, err := chromedpGenerator.RenderHTMLFile(tempFile.Name(), "main")
	assert.Nil(t, err)
	defer os.Remove(pdfFile)

	fi, err := os.Stat(pdfFile)
	assert.Nil(t, err)
	// Actually holds something
	assert.True(t, fi.Size() > 0)
}
