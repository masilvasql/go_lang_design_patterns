package project

import "fmt"

type RealVideoService struct{}

func (r *RealVideoService) PlayVideo(user string) {
	fmt.Printf("🎬 Reproduzindo vídeo para o usuário: %s\n", user)
}
