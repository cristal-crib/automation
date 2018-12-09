package sensors

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
	"go.uber.org/zap"

	"github.com/cristal-crib/spectrum/bridge/comms"
	"github.com/dghubble/sling"
)

var logger, _ = zap.NewProduction()

// GarageDoor represent a garage door
type GarageDoor struct {
	*accessory.Accessory
	Opener *service.GarageDoorOpener
	state  int
	guard  chan struct{}
	comms  *comms.HTTPComms
}

// NewGarageDoor create a garage door
func NewGarageDoor(info accessory.Info, comms *comms.HTTPComms) *GarageDoor {
	acc := GarageDoor{
		Accessory: accessory.New(info, accessory.TypeGarageDoorOpener),
		Opener:    service.NewGarageDoorOpener(),
		guard:     make(chan struct{}, 1),
		comms: comms,
	}
	acc.AddService(acc.Opener.Service)

	acc.Opener.CurrentDoorState.OnValueRemoteGet(acc.getState)
	acc.Opener.TargetDoorState.OnValueRemoteGet(acc.getTargetState)
	acc.Opener.TargetDoorState.OnValueRemoteUpdate(acc.setState)
	acc.Opener.CurrentDoorState.SetEventsEnabled(true)

	// Load a token into the guard channel
	acc.guard <- struct{}{}

	return &acc
}

func (d *GarageDoor) getTargetState() (state int) {
	switch d.getState() {
	case characteristic.CurrentDoorStateClosed, characteristic.CurrentDoorStateClosing:
		return characteristic.TargetDoorStateClosed
	case characteristic.CurrentDoorStateOpen, characteristic.CurrentDoorStateOpening:
		return characteristic.TargetDoorStateOpen
	default:
		return characteristic.CurrentDoorStateStopped
	}
}

func (d *GarageDoor) setState(to int) {
	req, err := sling.New().Get("http://localhost:30800/door-opener").Request()
	if err != nil {
		logger.Error("Unable to create request", zap.Error(err))
	}
	d.comms.Send(req)
	logger.Info("Door Opener state called")
}

func (d *GarageDoor) getState() (state int) {
	logger.Info("Get State called")
	// req, err := sling.New().Get("http://localhost:30800").Request()
	// d.comms.Do(req)
	return state
}
