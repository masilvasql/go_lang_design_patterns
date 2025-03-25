package dto

import "fmt"

type Pizza struct {
	Tamanho      string
	Massa        string
	Molho        string
	Ingredientes []string
}

func (p *Pizza) String() string {
	return fmt.Sprintf("Pizza %s, %s, %s, %v", p.Tamanho, p.Massa, p.Molho, p.Ingredientes)
}
