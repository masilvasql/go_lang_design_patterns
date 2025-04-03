package main

import (
	"github.com/masilvasql/prototype/client"
	"github.com/masilvasql/prototype/prototypes"
)

func main() {
	//creating original products
	laptop := &prototypes.Electronic{Brand: "Apple", Model: "MacBook Pro", Price: 1999.99}
	sofa := &prototypes.Furniture{Type: "Sofa", Color: "Blue", Price: 899.99}

	// Cloning and printing the products
	client.CloneAndPrint(laptop)
	client.CloneAndPrint(sofa)
}
