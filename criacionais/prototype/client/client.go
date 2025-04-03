package client

import (
	"fmt"
	"github.com/masilvasql/prototype/prototypes"
)

func CloneAndPrint(product prototypes.Product) {
	// Clone the product
	cloneProduct := product.Clone()

	fmt.Println("Produto original:", product.GetInfo())
	fmt.Println("Produto clonado: ", cloneProduct.GetInfo())
	fmt.Println("---")
}
