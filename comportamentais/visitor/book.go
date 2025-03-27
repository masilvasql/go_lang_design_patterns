package main

type Book struct {
	Price float64
	Name  string
}

func (b *Book) Accept(visitor Visitor) {
	visitor.VisitBook(b)
}

func (b *Book) GetPrice() float64 {
	return b.Price
}
