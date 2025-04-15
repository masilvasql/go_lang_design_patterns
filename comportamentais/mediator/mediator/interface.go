package mediator

type Participant interface {
	Receive(senderName, message string)
	GetName() string
}

type Mediator interface {
	SendMessage(sender Participant, message string)
	Register(user Participant)
}
