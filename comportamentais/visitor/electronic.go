package main

type Electronic struct {
	Price float64
	Brand string
}

func (e *Electronic) Accept(visitor Visitor) {
	visitor.VisitElectronic(e)
}

func (e *Electronic) GetPrice() float64 {
	return e.Price
}
