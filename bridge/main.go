package main

import (
	"github.com/cristal-crib/spectrum/bridge/comms"
	"github.com/cristal-crib/spectrum/bridge/sensors"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func main() {
	defer logger.Sync()
	infoG := accessory.Info{
		Name:         "CristalGarage",
		SerialNumber: "GDOOR-0001",
		Manufacturer: "CristalCrib",
		Model:        "CCGD",
	}
	comms := comms.NewHTTPComms()
	door := sensors.NewGarageDoor(infoG, comms)

	config := hc.Config{Pin: "00102009", StoragePath: "./tmp"}
	t, err := hc.NewIPTransport(config, door.Accessory)
	if err != nil {
		logger.Panic("Unable to start Homekit transport", zap.Error(err))
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
