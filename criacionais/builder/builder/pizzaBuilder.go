package builder

import (
	"github.com/masilvsql/builder/dto"
)

type PizzaBuilder struct {
	Pizza *dto.Pizza
}

func NewPizzaBuilder() Builder {
	return &PizzaBuilder{Pizza: &dto.Pizza{}}
}

func (pb *PizzaBuilder) SetTamanho(tamanho string) Builder {
	pb.Pizza.Tamanho = tamanho
	return pb
}

func (pb *PizzaBuilder) SetMassa(massa string) Builder {
	pb.Pizza.Massa = massa
	return pb
}

func (pb *PizzaBuilder) SetMolho(molho string) Builder {
	pb.Pizza.Molho = molho
	return pb
}

func (pb *PizzaBuilder) AddIngrediente(ingrediente string) Builder {
	pb.Pizza.Ingredientes = append(pb.Pizza.Ingredientes, ingrediente)
	return pb
}

func (pb *PizzaBuilder) Build() *dto.Pizza {
	return pb.Pizza
}
