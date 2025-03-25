package builder

import "github.com/masilvsql/builder/dto"

type Builder interface {
	SetTamanho(tamanho string) Builder
	SetMassa(massa string) Builder
	SetMolho(molho string) Builder
	AddIngrediente(ingrediente string) Builder
	Build() *dto.Pizza
}
