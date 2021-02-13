package pdf

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func RenderUrlToPdf(templateFile string, sel string) (string, error) {
	var pdfBuffer []byte

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Duration(config.Config.Timeout)*time.Second)
	defer cancel()

	absTemplateFile, err := filepath.Abs(templateFile)
	if err != nil {
		log.Print(err)
		return "", errors.New(fmt.Sprintf("Failed to create the temporary %v file"))
	}

	tasks := chromedp.Tasks{
		chromedp.Navigate("file://" + absTemplateFile),
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

	err = chromedp.Run(ctx, tasks)
	if err != nil {
		log.Println(err)
		return "", errors.New("Failed to render the PDF")
	}

	pdfFile := strings.Trim(absTemplateFile, ".html") + ".pdf"

	ioutil.WriteFile(pdfFile, pdfBuffer, 0644)
	if err != nil {
		log.Println(fmt.Sprintf("Error to write pdf"))
		return "", err
	}

	return pdfFile, err
}
