package main

import (
	"fmt"
	"github.com/masilvsql/builder/builder"
)

func main() {
	pizza1 := builder.NewPizzaBuilder().
		SetTamanho("Grande").
		SetMassa("Fina").
		SetMolho("Tomate").
		AddIngrediente("Queijo").
		AddIngrediente("Presunto").
		Build()

	fmt.Println(pizza1.String())

	pizza2 := builder.NewPizzaBuilder().
		SetTamanho("Pequena").
		SetMassa("Grossa").
		SetMolho("Barbecue").
		AddIngrediente("Queijo").
		AddIngrediente("Bacon").
		AddIngrediente("Cebola").
		Build()

	fmt.Println(pizza2.String())

}
