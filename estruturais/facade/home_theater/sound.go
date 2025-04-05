package home_theater

import "fmt"

type SoundSystem struct {
}

func (s *SoundSystem) TurnOn() {
	fmt.Println("Som Ligado")
}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Printf("Volume ajustado para %d.\n", level)
}

func (s *SoundSystem) TurnOff() {
	fmt.Println("Som desligado.")
}
