package sender

type CompressionDecorator struct {
	Sender MessageSender
}

func (c *CompressionDecorator) Send(msg string) string {
	// Compress message
	return "[Compactado] " + c.Sender.Send(msg)
}
