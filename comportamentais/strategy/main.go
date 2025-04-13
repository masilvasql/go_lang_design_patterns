package main

import (
	"fmt"
	"github.com/masilvasql/strategy/context"
	"github.com/masilvasql/strategy/strategy"
)

func main() {
	valorBase := 100.0
	calc := context.CalculadorFrete{}

	calc.SetStrategy(&strategy.FreteNormal{})
	fmt.Printf("Frete normal R$ %.2f\n", calc.CalcularFrete(valorBase))

	calc.SetStrategy(&strategy.FreteExpresso{})
	fmt.Printf("Frete expresso R$ %.2f\n", calc.CalcularFrete(valorBase))

	calc.SetStrategy(&strategy.FreteInternacional{})
	fmt.Printf("Frete internacional R$ %.2f\n", calc.CalcularFrete(valorBase))
}
