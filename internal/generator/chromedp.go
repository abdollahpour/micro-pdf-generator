package generator

import (
	"context"
	"errors"
	"io/ioutil"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/inspector"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

// ChromedpGenerator uses Chromedp and chromium to generate PDFs
type ChromedpGenerator struct {
	timeout int
	tempDir string
}

func NewChromedpGenerator(timeout int, tempDir string) *ChromedpGenerator {
	return &ChromedpGenerator{
		timeout: timeout,
		tempDir: tempDir,
	}
}

// RenderHTMLFile render PDF from html file
func (s ChromedpGenerator) RenderHTMLFile(htmlFile string, waitFor string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Duration(s.timeout)*time.Second)
	defer cancel()

	var pdfBuffer []byte

	tasks := chromedp.Tasks{
		network.Enable(),
		fetch.Enable(),
		inspector.Enable(),
		page.Enable(),
		chromedp.Navigate("file://" + htmlFile),
		chromedp.WaitVisible(waitFor, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			pdfBuffer = buf
			return nil
		}),
	}

	chromedp.ListenTarget(ctx, disableExternal(ctx, htmlFile))

	err := chromedp.Run(ctx, tasks)
	if err != nil {
		log.WithField("htmlFile", htmlFile).WithField("waitFor", waitFor).Warn("Failed to render the PDF")
		return "", errors.New("Failed to render the html file " + htmlFile)
	}

	pdfFile, err := ioutil.TempFile(s.tempDir, "*.pdf")
	if err != nil {
		log.Fatal("Failed to create temp file")
	}

	_, err = pdfFile.Write(pdfBuffer)
	if err != nil {
		log.Fatal("Failed to create temp file")
	}

	return pdfFile.Name(), nil
}

func disableExternal(ctx context.Context, htmlFile string) func(event interface{}) {
	return func(event interface{}) {
		switch ev := event.(type) {
		case *fetch.EventRequestPaused:
			go func() {
				c := chromedp.FromContext(ctx)
				ctx := cdp.WithExecutor(ctx, c.Target)
				if ev.Request.URL == "file://"+htmlFile {
					fetch.ContinueRequest(ev.RequestID).Do(ctx)
				} else {
					fetch.FailRequest(ev.RequestID, network.ErrorReasonAccessDenied).Do(ctx)
				}
			}()
		}
	}
}
