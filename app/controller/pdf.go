package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"genreport"
	"genreport/app/service"
)

// CreatePDF - GET /api/v1/create-pdf?orgID={orgID}&brandID={brandID}&reportID={reportID}&filename={filename}
func CreatePDF(c *gin.Context) {
	orgID := c.Query("orgID")
	orgID = strings.TrimSpace(orgID)
	if orgID == "" || orgID == "0" {
		c.String(
			http.StatusBadRequest,
			"organization ID is required",
		)
		return
	}

	brandID := c.Query("brandID")
	brandID = strings.TrimSpace(brandID)
	if brandID == "" || brandID == "0" {
		c.String(
			http.StatusBadRequest,
			"brand ID is required",
		)
		return
	}

	id := c.Query("reportID")
	id = strings.TrimSpace(id)
	if id == "" {
		c.String(
			http.StatusBadRequest,
			"report ID is required",
		)
		return
	}

	filename := c.Query("filename")
	filename = strings.TrimSpace(filename)
	if filename == "" {
		c.String(
			http.StatusBadRequest,
			"filename is required",
		)
		return
	}
	if !strings.HasSuffix(filename, ".pdf") {
		filename += ".pdf"
	}

	// http://ip:port?report-id={id}
	inputURL := genreport.GetURLReportUI() + "?report-id=" + id
	filepath := genreport.GetDirReport() + orgID + "/" + brandID + "/"

	err := service.CreatePDF(genreport.Browser, inputURL, filepath, filename)
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.String(
		http.StatusOK,
		filename,
	)
}
