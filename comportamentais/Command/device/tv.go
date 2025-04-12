package device

import "fmt"

type TV struct{}

func (t *TV) On() {
	fmt.Println("TV On")
}
