package main

import "github.com/masilvasql/flyweight/project"

func main() {
	forest := &project.Forest{}

	//planta várias árvores com as mesmas prpriedades compartilhadas
	forest.PlantTree(10, 20, "Oak", "Green", "Rough")
	forest.PlantTree(30, 40, "Oak", "Green", "Rough")
	forest.PlantTree(50, 60, "Pine", "Dark Green", "Smooth")
	forest.PlantTree(70, 80, "Oak", "Green", "Rough")

	forest.Draw()
}
