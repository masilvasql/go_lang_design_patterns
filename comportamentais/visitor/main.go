package main

import "fmt"

func main() {
	elements := []Element{
		&Book{Price: 20, Name: "Book 1"},
		&Electronic{Price: 100, Brand: "Brand 1"},
		&Book{Price: 30, Name: "Book 2"},
	}

	discountVisitor := &DiscountVisitor{}

	fmt.Println("Calculando descontos:")
	for _, element := range elements {
		element.Accept(discountVisitor)
	}

	fmt.Printf("Desconto total: %.2f\n", discountVisitor.TotalDiscount)
}
