package interfaces

// Observer define a interface para os objetos que serão notificados
type Observer interface {
	Update(message string)
}

// Subject define a interface para o objeto que mantém os observadores
type Subject interface {
	Register(o Observer)
	Remove(o Observer)
	NotifyObservers()
}
