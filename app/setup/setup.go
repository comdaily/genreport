// Package setup - set up the application
package setup

import (
	"time"

	gconfig "github.com/pilinux/gorest/config"

	"genreport"
)

// EnvConfig - set environment variables
func EnvConfig(tNow time.Time) error {
	// load configs
	err := gconfig.Config()
	if err != nil {
		return err
	}

	// read configs
	config := gconfig.GetConfig()

	// set version
	genreport.Version = config.Version
	// set environment
	genreport.Env = config.Server.ServerEnv
	// set server start time
	genreport.StartTime = tNow

	return nil
}

// SetPath - set paths to directories
func SetPath() error {
	// set reporting directory
	err := genreport.SetDirReport()
	if err != nil {
		return err
	}

	return nil
}

// SetRemoteURLs - set remote URLs
func SetRemoteURLs() error {
	// set remote URLs
	err := genreport.SetRemoteURLs()
	if err != nil {
		return err
	}

	return nil
}
