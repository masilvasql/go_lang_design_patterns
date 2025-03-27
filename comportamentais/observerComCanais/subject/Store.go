package subject

type Store struct {
	Observers []chan string
	Promotion string
}

func (s *Store) Register(o chan string) {
	s.Observers = append(s.Observers, o)
}

func (s *Store) Remove(o chan string) {
	for i, observer := range s.Observers {
		if observer == o {
			s.Observers = append(s.Observers[:i], s.Observers[i+1:]...)
			break
		}
	}
}

func (s *Store) NotifyObservers() {
	for _, observer := range s.Observers {
		observer <- s.Promotion
	}
}

func (s *Store) SetPromotion(promo string) {
	s.Promotion = promo
	s.NotifyObservers()
}
