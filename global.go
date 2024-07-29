// Package genreport : This package is used to generate the
// full report for the given data and save it in PDF format.
package genreport

import (
	"time"

	"github.com/go-rod/rod"
)

// Browser is a global browser instance
var Browser *rod.Browser

// global variables
var (
	// released version or commit number
	Version string
	// environment
	Env string
	// server start time
	StartTime time.Time
)
