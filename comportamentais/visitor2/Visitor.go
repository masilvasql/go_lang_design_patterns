package main

type Visitor interface {
	VisitPessoa(p *Pessoa)
}
