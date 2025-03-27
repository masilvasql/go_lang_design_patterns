package main

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Accept(visitor Visitor) {
	visitor.VisitPessoa(p)
}

func (p *Pessoa) GetNome() string {
	return p.Nome
}

func (p *Pessoa) GetIdade() int {
	return p.Idade
}
