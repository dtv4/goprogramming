package base

type Subject struct {
	state     int
	observers []Observer
}

func (s *Subject) attach(o Observer) {
	s.observers = append(s.observers, o)
}
func (s *Subject) setState(i int) {
	s.state = i
	s.notifyAllObservers()
}
func (s *Subject) notifyAllObservers() {
	for _, obs := range s.observers {
		obs.update()
	}
}
func (s *Subject) getState() int {
	return s.state
}
