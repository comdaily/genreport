package genreport

import (
	"fmt"
	"os"
	"strings"
)

// directories
var (
	DirReport string
)

// SetDirReport sets the reporting directory path
func SetDirReport() error {
	DirReport = strings.TrimSpace(os.Getenv("DIR_REPORT"))
	if DirReport == "" {
		return fmt.Errorf("DIR_REPORT not set")
	}

	// add trailing slash if missing
	if !strings.HasSuffix(DirReport, "/") {
		DirReport += "/"
	}

	return nil
}

// GetDirReport returns the reporting directory path
func GetDirReport() string {
	return DirReport
}
