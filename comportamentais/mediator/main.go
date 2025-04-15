package main

import (
	"github.com/masilvasql/mediator/chatroom"
	"github.com/masilvasql/mediator/user"
)

func main() {

	chat := chatroom.NewChatRoom()

	marcelo := user.NewUser("Marcelo", chat)
	rafael := user.NewUser("Rafael", chat)
	joao := user.NewUser("João", chat)

	chat.Register(marcelo)
	chat.Register(rafael)
	chat.Register(joao)

	marcelo.Send("Oi, tudo bem?")
	rafael.Send("Oi, tudo bem? E você?")
	joao.Send("Oi, tudo bem? E você? E o trabalho?")

}
