// Package router provides the core routing functionality
package router

import (
	"github.com/gin-gonic/gin"
	gconfig "github.com/pilinux/gorest/config"
	glib "github.com/pilinux/gorest/lib"
	gmiddleware "github.com/pilinux/gorest/lib/middleware"

	"genreport/app/controller"
)

// SetupRouter sets up the router
func SetupRouter(configure *gconfig.Configuration) (*gin.Engine, error) {
	if gconfig.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin.Default() = gin.New() + gin.Logger() + gin.Recovery()
	r := gin.Default()

	// Which proxy to trust:
	// disable this feature as it still fails
	// to provide the real client IP in
	// different scenarios
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return r, err
	}

	// when using Cloudflare's CDN:
	// router.TrustedPlatform = gin.PlatformCloudflare
	//
	// when running on Google App Engine:
	// router.TrustedPlatform = gin.PlatformGoogleAppEngine
	//
	/*
		when using apache or nginx reverse proxy
		without Cloudflare's CDN or Google App Engine

		config for nginx:
		=================
		proxy_set_header X-Real-IP       $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	*/
	// router.TrustedPlatform = "X-Real-Ip"
	//
	// set TrustedPlatform to get the real client IP
	trustedPlatform := configure.Security.TrustedPlatform
	if trustedPlatform == "cf" {
		trustedPlatform = gin.PlatformCloudflare
	}
	if trustedPlatform == "google" {
		trustedPlatform = gin.PlatformGoogleAppEngine
	}
	r.TrustedPlatform = trustedPlatform

	// CORS
	if gconfig.IsCORS() {
		r.Use(gmiddleware.CORS(configure.Security.CORS))
	}

	// Origin
	if gconfig.IsOriginCheck() {
		r.Use(gmiddleware.CheckOrigin())
	}

	// Sentry.io
	if gconfig.IsSentry() {
		r.Use(gmiddleware.SentryCapture(
			configure.Logger.SentryDsn,
			configure.Server.ServerEnv,
			configure.Version,
		))
	}

	// WAF
	if gconfig.IsWAF() {
		r.Use(gmiddleware.Firewall(
			configure.Security.Firewall.ListType,
			configure.Security.Firewall.IP,
		))
	}

	// Rate Limiter
	if gconfig.IsRateLimit() {
		// create a rate limiter instance
		limiterInstance, err := glib.InitRateLimiter(
			configure.Security.RateLimit,
			trustedPlatform,
		)
		if err != nil {
			return r, err
		}
		r.Use(gmiddleware.RateLimit(limiterInstance))
	}

	// API Status
	r.GET("", controller.APIStatus)

	// API:v1
	v1 := r.Group("/api/v1/")
	{
		// create PDF report
		v1.GET("create-pdf", controller.CreatePDF)
	}

	return r, nil
}
