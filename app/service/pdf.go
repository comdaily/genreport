// Package service ...
package service

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// CreatePDF fetches the HTML page at the URL given and
// creates a PDF file in the specified directory
func CreatePDF(browser *rod.Browser, inputURL, path, filename string) error {
	// create a new page and navigate to the URL
	page := browser.MustPage(inputURL)
	defer page.MustClose()

	// set a timeout of 60 seconds and wait for the page to load
	page.Timeout(60 * time.Second).MustWaitLoad()

	// start to analyze request events
	wait := page.MustWaitRequestIdle()

	// wait until the page is idle
	wait()

	a4Width := 8.27
	a4Height := 11.7
	a4MarginTop := 0.0
	a4MarginBottom := 0.0
	a4MarginLeft := 0.0
	a4MarginRight := 0.0
	p := proto.PagePrintToPDF{
		PrintBackground: true,
		PaperWidth:      &a4Width,
		PaperHeight:     &a4Height,
		MarginTop:       &a4MarginTop,
		MarginBottom:    &a4MarginBottom,
		MarginLeft:      &a4MarginLeft,
		MarginRight:     &a4MarginRight,
	}

	// print page as PDF
	// page.MustPDF("filename.pdf")
	r, e := page.PDF(&p)
	if e != nil {
		return e
	}
	bin, e := io.ReadAll(r)
	if e != nil {
		return e
	}
	// size of PDF must be greater than 1000 bytes
	if len(bin) < 1000 {
		return fmt.Errorf("PDF size is less than 1000 bytes")
	}

	// create the directories if they do not exist
	dir := filepath.Dir(path)
	e = os.MkdirAll(dir, 0755)
	if e != nil {
		return e
	}

	// save the PDF to a file in the proper working directory
	file, e := os.Create(path + filename)
	if e != nil {
		return e
	}
	defer file.Close()

	_, e = file.Write(bin)
	if e != nil {
		return e
	}

	// optionally, flush the buffer to ensure data is written immediately
	e = file.Sync()
	if e != nil {
		return e
	}

	return nil
}
