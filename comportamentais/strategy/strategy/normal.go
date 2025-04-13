package strategy

type FreteNormal struct{}

func (f *FreteNormal) Calcular(valor float64) float64 {
	return valor + 10
}
