package composite

import (
	"fmt"
	"github.com/masilvasql/composite/interfaces"
)

type Folder struct {
	Name     string
	children []interfaces.FileSystemComponent
}

func (f *Folder) Add(child interfaces.FileSystemComponent) {
	f.children = append(f.children, child)
}

func (f *Folder) Show(indent string) {
	fmt.Println(indent + "[Folder] " + f.Name)
	for _, child := range f.children {
		child.Show(indent + "  ")
	}
}
