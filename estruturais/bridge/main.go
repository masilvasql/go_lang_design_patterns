package main

import (
	"github.com/masilvasql/bridge/device"
	remote2 "github.com/masilvasql/bridge/remote"
)

func main() {
	tv := &device.TV{}
	tvRemote := &remote2.Remote{Device: tv}

	tvRemote.TogglePower()
	tvRemote.VolumeUp()
	tvRemote.VolumeUp()
	tvRemote.VolumeDown()

	advRemote := &remote2.AdvancedRemote{Remote: tvRemote}
	advRemote.Mute()

	radio := &device.Radio{}
	radioRemote := &remote2.Remote{Device: radio}
	advRemoteRadio := &remote2.AdvancedRemote{Remote: radioRemote}

	advRemoteRadio.TogglePower()
	advRemoteRadio.VolumeUp()
	advRemoteRadio.Mute()
}
