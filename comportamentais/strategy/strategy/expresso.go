package strategy

type FreteExpresso struct{}

func (f *FreteExpresso) Calcular(valor float64) float64 {
	return valor + 30
}
