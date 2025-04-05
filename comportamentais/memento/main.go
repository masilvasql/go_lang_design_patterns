package main

import (
	"github.com/masilvasql/memento/caretaker"
	"github.com/masilvasql/memento/originator"
)

func main() {
	heroi := &originator.Heroi{}
	arquivista := &caretaker.Arquivista{}

	// Aprendendo habilidades
	heroi.AprenderHabilidade("Voo")
	arquivista.Guardar(heroi.SalvarHabilidade())

	heroi.AprenderHabilidade("Super Força")
	arquivista.Guardar(heroi.SalvarHabilidade())

	heroi.AprenderHabilidade("Invisibilidade")
	arquivista.Guardar(heroi.SalvarHabilidade())

	// Restaurando habilidades
	heroi.RestaurarHabilidade(arquivista.Recuperar(0))
}
