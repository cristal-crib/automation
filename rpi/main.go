package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var rel *gpio.RelayDriver
var reedBottom *gpio.ButtonDriver
var reedTop *gpio.ButtonDriver

func main() {
	r := raspi.NewAdaptor()
	rel = gpio.NewRelayDriver(r, "7")
	reedBottom = gpio.NewButtonDriver(r, "11")
	reedTop = gpio.NewButtonDriver(r, "13")
	rel = gpio.NewRelayDriver(r, "7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			rel.On()
			time.Sleep(500 * time.Millisecond)
			rel.Off()
		})
	}

	robot := gobot.NewRobot("Spectrum Garage",
		[]gobot.Connection{r},
		[]gobot.Device{rel},
		work,
	)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/door-opener", doorOpener)
	e.GET("/door-state", doorState)

	// Start server
	e.Logger.Fatal(e.Start(":30800"))

	robot.Start()
}

func doorOpener(c echo.Context) error {
	rel.On()
	time.Sleep(200 * time.Millisecond)
	rel.Off()
	return c.String(http.StatusOK, "Relay command received")
}

func doorState(c echo.Context) error {

	return c.String(http.StatusOK, "")
}
