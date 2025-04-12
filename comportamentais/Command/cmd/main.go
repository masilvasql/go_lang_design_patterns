package main

import (
	"github.com/masilvasql/command/command"
	"github.com/masilvasql/command/device"
	"github.com/masilvasql/command/invoker"
)

func main() {
	//devices
	light := &device.Light{}
	tv := &device.TV{}

	//commands
	lightOn := &command.LightOnCommand{Light: light}
	lightOff := &command.LightOffCommand{Light: light}
	tvOn := &command.TVOnCommand{TV: tv}

	//invoker
	remote := invoker.NewRemoteControl()
	remote.SetCommand("light_on", lightOn)
	remote.SetCommand("light_off", lightOff)
	remote.SetCommand("tv_on", tvOn)

	//use
	remote.PressButton("light_on")
	remote.PressButton("light_off")
	remote.PressButton("tv_on")
}
