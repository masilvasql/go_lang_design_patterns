package command

import (
	"github.com/masilvasql/command/device"
)

type LightOffCommand struct {
	Light *device.Light
}

func (c *LightOffCommand) Execute() {
	c.Light.Off()
}
