package main

type Visitor interface {
	VisitElectronic(e *Electronic)
	VisitBook(b *Book)
}
