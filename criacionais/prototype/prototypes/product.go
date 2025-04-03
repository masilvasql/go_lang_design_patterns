package prototypes

// Product is a prototype interface that defines the Clone method.
type Product interface {
	Clone() Product
	GetInfo() string
}
