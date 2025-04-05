package remote

import "github.com/masilvasql/bridge/device"

type Remote struct {
	Device device.Device
}

func (r *Remote) TogglePower() {
	if r.Device.IsEnabled() {
		r.Device.Disable()
	} else {
		r.Device.Enable()
	}
}

func (r *Remote) VolumeDown() {
	vol := r.Device.GetVolume()
	r.Device.SetVolume(vol - 1)
}

func (r *Remote) VolumeUp() {
	vol := r.Device.GetVolume()
	r.Device.SetVolume(vol + 1)
}
