package originator

import (
	"fmt"
	"github.com/masilvasql/memento/memento"
)

type Heroi struct {
	Habilidades string
}

func (h *Heroi) AprenderHabilidade(nova string) {
	h.Habilidades = nova
	fmt.Println("Kael aprendeu:", h.Habilidades)
}

func (h *Heroi) SalvarHabilidade() memento.Memento {
	return memento.Memento{h.Habilidades}
}

func (h *Heroi) RestaurarHabilidade(m memento.Memento) {
	h.Habilidades = m.Habilidades
	fmt.Println("Kael recuperou:", h.Habilidades)
}
