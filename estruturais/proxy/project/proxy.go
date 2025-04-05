package project

import "fmt"

type ProxyVideoService struct {
	realService  *RealVideoService
	allowedUsers map[string]bool
}

func NewProxyVideoService(allowedUsers map[string]bool) *ProxyVideoService {
	return &ProxyVideoService{
		realService:  &RealVideoService{},
		allowedUsers: allowedUsers,
	}
}

func (p *ProxyVideoService) PlayVideo(user string) {
	if !p.allowedUsers[user] {
		fmt.Printf("❌ Acesso negado para o usuário: %s\n", user)
		return
	}

	fmt.Printf("✅ Acesso concedido para o usuário: %s\n", user)
	p.realService.PlayVideo(user)

}
