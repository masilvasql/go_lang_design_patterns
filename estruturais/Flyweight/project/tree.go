package project

import "fmt"

// Flyweight - dados imutáveis
type TreeType struct {
	Name    string
	Color   string
	Texture string
}

func (t *TreeType) Draw(x int, y int) {
	fmt.Printf("Drawing tree of type %s at coordinates (%d, %d) with color %s and texture %s\n", t.Name, x, y, t.Color, t.Texture)
}

// Contexto - contem os dados mutáveis
type Tree struct {
	X        int
	Y        int
	TreeType *TreeType
}

func (t *Tree) Draw() {
	t.TreeType.Draw(t.X, t.Y)
}
