package project

var treeTypes = make(map[string]*TreeType)

// GetTreeType - retorna o tipo de árvore correspondente ao nome
func GetTreeType(name string, color string, texture string) *TreeType {
	key := name + "_" + color + "_" + texture
	if treeType, exists := treeTypes[key]; exists {
		return treeType
	}

	treeType := &TreeType{
		Name:    name,
		Color:   color,
		Texture: texture,
	}
	treeTypes[key] = treeType
	return treeType
}
