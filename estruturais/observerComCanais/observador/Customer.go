package observador

import "fmt"

type Customer struct {
	Name string
}

func (c *Customer) Update(promotion string) {
	fmt.Printf("Customer %s received promotion: %s\n", c.Name, promotion)
}
