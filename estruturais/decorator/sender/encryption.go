package sender

// EncryptionDecorator is a struct that implements the MessageSender interface
type EncryptionDecorator struct {
	Sender MessageSender
}

func (e *EncryptionDecorator) Send(msg string) string {
	// Encrypt message
	return "[Criptografado] " + e.Sender.Send(msg)
}
