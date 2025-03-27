package observer

import "fmt"

// Customer é um observador que reage às promoções
type Customer struct {
	Name string
}

// Update é o método chamado pelo sujeito para notificar o cliente
func (c *Customer) Update(message string) {
	fmt.Printf("%s recebeu a promoção: %s\n", c.Name, message)
}
