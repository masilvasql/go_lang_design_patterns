package main

import (
	"fmt"
	"github.com/masilvasql/decorator/sender"
)

func main() {
	basicSender := &sender.BasicSender{}
	basic := basicSender.Send("Hello World")

	encrptionSender := &sender.EncryptionDecorator{Sender: basicSender}
	encrypted := encrptionSender.Send("Hello World")

	compressionSender := &sender.CompressionDecorator{Sender: basicSender}
	compressed := compressionSender.Send("Hello World")

	fullDecorator := &sender.CompressionDecorator{Sender: &sender.EncryptionDecorator{Sender: basicSender}}
	full := fullDecorator.Send("Hello World")

	fmt.Println(basic)
	fmt.Println("-----")
	fmt.Println(encrypted)
	fmt.Println("-----")
	fmt.Println(compressed)
	fmt.Println("-----")
	fmt.Println(full)

}
