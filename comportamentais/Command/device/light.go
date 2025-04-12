package device

import "fmt"

type Light struct{}

func (l *Light) On() {
	fmt.Println("Light On")
}

func (l *Light) Off() {
	fmt.Println("Light Off")
}
