package home_theater

import "fmt"

type Light struct {
}

func (l *Light) DimLight() {
	fmt.Println("Dim light")
}

func (l *Light) TurnOn() {
	fmt.Println("Light on")
}

func (l *Light) TurnOff() {
	fmt.Println("Light off")
}
