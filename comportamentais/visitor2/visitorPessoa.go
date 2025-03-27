package main

type VisitorPessoa struct {
	isLegalAge bool
}

func (vp *VisitorPessoa) VisitPessoa(p *Pessoa) {
	if p.Idade >= 18 {
		vp.isLegalAge = true
	}
}
