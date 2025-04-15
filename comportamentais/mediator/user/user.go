package user

import "github.com/masilvasql/mediator/mediator"

type User struct {
	Name     string
	ChatRoom mediator.Mediator
}

func NewUser(name string, m mediator.Mediator) *User {
	return &User{
		Name:     name,
		ChatRoom: m,
	}
}

func (u *User) Send(message string) {
	u.ChatRoom.SendMessage(u, message)
}

func (u *User) Receive(senderName, message string) {
	println("[" + senderName + " → " + u.Name + "]: " + message)
}

func (u *User) GetName() string {
	return u.Name
}
