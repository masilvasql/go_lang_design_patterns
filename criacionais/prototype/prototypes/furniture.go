package prototypes

import "fmt"

// Furniture is a concrete prototype that implements the Product interface.
type Furniture struct {
	Type  string
	Color string
	Price float64
}

func (f *Furniture) Clone() Product {
	return &Furniture{
		Type:  f.Type,
		Color: f.Color,
		Price: f.Price,
	}
}

func (f *Furniture) GetInfo() string {
	return fmt.Sprintf("Móvel: %s - Cor: %s - R$ %.2f", f.Type, f.Color, f.Price)
}
