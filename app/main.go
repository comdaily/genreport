// Package main ...
package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	gconfig "github.com/pilinux/gorest/config"

	"genreport"
	"genreport/app/router"
	"genreport/app/setup"
)

func main() {
	timeNow := time.Now()
	fmt.Println("Starting user API server at:", timeNow)

	// set configs
	err := setup.EnvConfig(timeNow)
	if err != nil {
		fmt.Println(err)
		return
	}

	// set up paths to directories
	err = setup.SetPath()
	if err != nil {
		fmt.Println(err)
		return
	}

	// set remote URLs
	err = setup.SetRemoteURLs()
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a new browser instance
	genreport.Browser = rod.New().MustConnect()
	defer genreport.Browser.MustClose()

	// read configs
	config := gconfig.GetConfig()

	// setup router
	r, err := router.SetupRouter(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// start server
	err = r.Run(config.Server.ServerHost + ":" + config.Server.ServerPort)
	if err != nil {
		fmt.Println(err)
		return
	}
}
