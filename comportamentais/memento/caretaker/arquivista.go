package caretaker

import "github.com/masilvasql/memento/memento"

type Arquivista struct {
	historico []memento.Memento
}

func (a *Arquivista) Guardar(m memento.Memento) {
	a.historico = append(a.historico, m)
}

func (a *Arquivista) Recuperar(index int) memento.Memento {
	return a.historico[index]
}
