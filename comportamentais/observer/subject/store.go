package subject

import "github.com/masilvsql/observer/interfaces"

// Store é o sujeito que mantém promoções e notifica observadores
type Store struct {
	Observers []interfaces.Observer
	Promotion string
}

// RegisterObserver adiciona um observador à lista
func (s *Store) Register(o interfaces.Observer) {
	s.Observers = append(s.Observers, o)
}

// RemoveObserver remove um observador da lista
func (s *Store) Remove(o interfaces.Observer) {
	for i, observer := range s.Observers {
		if observer == o {
			s.Observers = append(s.Observers[:i], s.Observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers envia a promoção atual para todos os observadores
func (s *Store) NotifyObservers() {
	for _, observer := range s.Observers {
		observer.Update(s.Promotion)
	}
}

// SetPromotion atualiza a promoção e notifica os observadores
func (s *Store) SetPromotion(promo string) {
	s.Promotion = promo
	s.NotifyObservers()
}
