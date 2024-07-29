package genreport

import (
	"fmt"
	"os"
	"strings"
)

// remote URLs
var (
	ReportUI string
)

// SetRemoteURLs sets the remote URLs
func SetRemoteURLs() error {
	ReportUI = strings.TrimSpace(os.Getenv("REPORT_UI"))
	if ReportUI == "" {
		return fmt.Errorf("REPORT_UI not set")
	}

	return nil
}

// GetURLReportUI returns the remote URL for the report UI
func GetURLReportUI() string {
	return ReportUI
}
