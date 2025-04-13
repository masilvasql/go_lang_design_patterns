package context

import "github.com/masilvasql/strategy/strategy"

type CalculadorFrete struct {
	Strategy strategy.FreteStrategy
}

func (c *CalculadorFrete) CalcularFrete(valor float64) float64 {
	return c.Strategy.Calcular(valor)
}

func (c *CalculadorFrete) SetStrategy(strategy strategy.FreteStrategy) {
	c.Strategy = strategy
}
