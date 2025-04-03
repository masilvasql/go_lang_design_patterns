package sender

// MessageSender is an interface that defines the Send method
type MessageSender interface {
	Send(msg string) string
}

// BasicSender is a struct that implements the MessageSender interface
type BasicSender struct{}

func (b *BasicSender) Send(msg string) string {
	// Send message
	return "Enviando mensagem : " + msg
}
