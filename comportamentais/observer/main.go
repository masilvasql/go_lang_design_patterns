package main

import (
	"github.com/masilvsql/observer/observer"
	"github.com/masilvsql/observer/subject"
)

func main() {

	store := &subject.Store{}

	customer1 := &observer.Customer{Name: "Alice"}
	customer2 := &observer.Customer{Name: "Bob"}
	customer3 := &observer.Customer{Name: "Charlie"}

	store.Register(customer1)
	store.Register(customer2)
	store.Register(customer3)

	store.SetPromotion("50% de desconto em todos os produtos!")

	store.Remove(customer2)

	store.SetPromotion("Frete grátis em compras acima de R$ 100,00!")

}
