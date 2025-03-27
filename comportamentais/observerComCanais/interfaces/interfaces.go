package interfaces

type ObserverComCanais interface {
	Register(ch chan string)
	Remove(ch chan string)
	NotifyObservers()
}
