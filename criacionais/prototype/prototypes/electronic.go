package prototypes

import "fmt"

// Electronic is a concrete prototype that implements the Product interface.
type Electronic struct {
	Brand string
	Model string
	Price float64
}

func (e *Electronic) Clone() Product {
	return &Electronic{
		Brand: e.Brand,
		Model: e.Model,
		Price: e.Price,
	}
}

func (e *Electronic) GetInfo() string {
	return fmt.Sprintf("Eletrônico: %s %s - R$ %.2f", e.Brand, e.Model, e.Price)
}
