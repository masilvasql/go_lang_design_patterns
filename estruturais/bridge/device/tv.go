package device

import "fmt"

type TV struct {
	enabled bool
	volume  int
}

func (t *TV) IsEnabled() bool {
	return t.enabled
}

func (t *TV) Enable() {
	t.enabled = true
	fmt.Println("TV is enabled")
}

func (t *TV) Disable() {
	t.enabled = false
	fmt.Println("TV is disabled")
}

func (t *TV) GetVolume() int {
	return t.volume
}

func (t *TV) SetVolume(volume int) {
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}
	t.volume = volume
	fmt.Printf("TV volume set to %d\n", t.volume)
}
