package pdf

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Generator interface
type Generator interface {
	RenderHTMLFile(templateFile string, sel string) (string, error)
}

// ChromedpGenerator uses Chromedp and chromium to generate PDFs
type ChromedpGenerator struct {
	Config config.Configuration
}

// RenderHTMLFile render PDF from html file
func (s ChromedpGenerator) RenderHTMLFile(htmlFile string, sel string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Duration(s.Config.Timeout)*time.Second)
	defer cancel()

	var pdfBuffer []byte

	tasks := chromedp.Tasks{
		chromedp.Navigate("file://" + path.Join(s.Config.TempDir, htmlFile)),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			pdfBuffer = buf
			return nil
		}),
	}

	err := chromedp.Run(ctx, tasks)
	if err != nil {
		log.Println(err)
		return "", errors.New("Failed to render the PDF")
	}

	pdfFile, err := ioutil.TempFile("", htmlFile)
	if err != nil {
		log.Println(err)
		return "", errors.New("Failed to render the PDF")
	}

	pdfFile.Write(pdfBuffer)
	if err != nil {
		log.Println(fmt.Sprintf("Error to write pdf"))
		return "", err
	}

	return pdfFile.Name(), err
}
