package project

type Forest struct {
	Trees []*Tree
}

func (f *Forest) PlantTree(x int, y int, name string, color string, texture string) {
	treeType := GetTreeType(name, color, texture)
	tree := &Tree{
		X:        x,
		Y:        y,
		TreeType: treeType,
	}
	f.Trees = append(f.Trees, tree)
}

func (f *Forest) Draw() {
	for _, tree := range f.Trees {
		tree.Draw()
	}
}
