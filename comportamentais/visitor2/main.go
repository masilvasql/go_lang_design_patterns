package main

import "fmt"

func main() {
	elements := []Element{
		&Pessoa{Idade: 17},
		&Pessoa{Idade: 18},
		&Pessoa{Idade: 19},
		&Pessoa{Idade: 15},
	}

	visitor := &VisitorPessoa{}
	for _, e := range elements {
		e.Accept(visitor)
		if visitor.isLegalAge {
			fmt.Println("Legal age")
		} else {
			fmt.Println("Not legal age")
		}
	}
}
