// Package controller ...
package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	grenderer "github.com/pilinux/gorest/lib/renderer"

	"genreport"
)

// AppStatus - status of the API
type AppStatus struct {
	Status    string    `json:"status"`
	Version   string    `json:"version"`
	Env       string    `json:"env"`
	StartTime time.Time `json:"startTime"`
	Uptime    string    `json:"uptime"`
}

// APIStatus - return status of the API
func APIStatus(c *gin.Context) {
	appStatus := AppStatus{}
	appStatus.Status = "ok"
	appStatus.Version = genreport.Version
	appStatus.Env = genreport.Env
	appStatus.StartTime = genreport.StartTime
	appStatus.Uptime = time.Since(genreport.StartTime).String()

	grenderer.Render(c, appStatus, http.StatusOK)
}
