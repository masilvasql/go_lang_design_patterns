package main

import "fmt"

type DiscountVisitor struct {
	TotalDiscount float64
}

func (v *DiscountVisitor) VisitBook(book *Book) {
	discount := book.GetPrice() * 0.10 // 10% de desconto para livros
	v.TotalDiscount += discount
	fmt.Printf("Livro: %s - Desconto: %.2f\n", book.Name, discount)
}

func (v *DiscountVisitor) VisitElectronic(electronic *Electronic) {
	discount := electronic.GetPrice() * 0.05 // 5% de desconto para eletrônicos
	v.TotalDiscount += discount
	fmt.Printf("Eletrônico: %s - Desconto: %.2f\n", electronic.Brand, discount)
}
