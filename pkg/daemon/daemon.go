package daemon

import (
	"net/http"

	"github.com/hahwul/jqueen/pkg/daemon"
	job "github.com/hahwul/jqueen/pkg/job"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RunDaemon is start rest api server
func RunDaemon() {
	var queue *job.JobQueue = job.New()
	_= queue
	e := echo.New()
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "jqueen")
	})

	e.GET("/add", func(c echo.Context) error {
		return daemon.AddJob(c, queue)		
	})



	e.Logger.Fatal(e.Start(":6886")) 
}
