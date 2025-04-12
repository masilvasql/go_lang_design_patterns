package command

import "github.com/masilvasql/command/device"

type TVOnCommand struct {
	TV *device.TV
}

func (c *TVOnCommand) Execute() {
	c.TV.On()
}
