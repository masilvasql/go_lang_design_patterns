package home_theater

import "fmt"

type Screen struct {
}

func (s *Screen) Lower() {
	fmt.Println("Tela baixada.")
}

func (s *Screen) Raise() {
	fmt.Println("Tela levantada.")
}
