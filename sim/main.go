package main

import (
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func main() {
	logger.Info("Starting Spectrum IOT Simulator")
	e := echo.New()
	e.Group("/garage")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	err := e.Start(":30800")
	if err != nil {
		logger.Fatal("Unable to start server", zap.Error(err))
	}
}
