package chatroom

import "github.com/masilvasql/mediator/mediator"

type ChatRoom struct {
	users []mediator.Participant
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (c *ChatRoom) Register(u mediator.Participant) {
	c.users = append(c.users, u)
}

func (c *ChatRoom) SendMessage(sender mediator.Participant, message string) {
	for _, u := range c.users {
		if u.GetName() != sender.GetName() {
			u.Receive(sender.GetName(), message)
		}
	}
}
