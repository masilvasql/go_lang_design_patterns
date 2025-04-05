package home_theater

type HomeTheaterFacade struct {
	light       *Light
	soundSystem *SoundSystem
	screen      *Screen
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		light:       &Light{},
		screen:      &Screen{},
		soundSystem: &SoundSystem{},
	}
}

func (f *HomeTheaterFacade) StartMoviewNight() {
	f.light.DimLight()
	f.screen.Lower()
	f.soundSystem.TurnOn()
	f.soundSystem.SetVolume(8)
}

func (f *HomeTheaterFacade) EndMovieNight() {
	f.soundSystem.TurnOff()
	f.screen.Raise()
	f.light.TurnOn()
}
