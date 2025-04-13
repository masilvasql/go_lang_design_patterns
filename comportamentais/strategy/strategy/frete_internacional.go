package strategy

type FreteInternacional struct{}

func (f *FreteInternacional) Calcular(valor float64) float64 {
	return valor + 60
}
