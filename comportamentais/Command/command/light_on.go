package command

import (
	"github.com/masilvasql/command/device"
)

type LightOnCommand struct {
	Light *device.Light
}

func (c *LightOnCommand) Execute() {
	c.Light.On()
}
