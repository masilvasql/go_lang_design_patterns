package remote

type AdvancedRemote struct {
	*Remote
}

func (a *AdvancedRemote) Mute() {
	a.Device.SetVolume(0)
}
