package device

import "fmt"

type Radio struct {
	enabled bool
	volume  int
}

func (r Radio) IsEnabled() bool {
	return r.enabled
}

func (r Radio) Enable() {
	r.enabled = true
	fmt.Println("Radio is enabled")
}

func (r Radio) Disable() {
	r.enabled = false
	fmt.Println("Radio is disabled")
}

func (r Radio) GetVolume() int {
	return r.volume
}

func (r Radio) SetVolume(volume int) {
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}
	r.volume = volume
	fmt.Printf("Radio volume set to %d\n", r.volume)
}
