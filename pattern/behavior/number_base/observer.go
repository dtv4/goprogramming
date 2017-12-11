package base

import "fmt"
import "strconv"

type Observer interface {
	update()
}

type BinaryObserver struct {
	subject *Subject
}

func (bo BinaryObserver) update() {
	fmt.Println("Binary", strconv.FormatInt(int64(bo.subject.getState()), 2))
}

type HexaObserver struct {
	subject *Subject
}

func (ho HexaObserver) update() {
	fmt.Println("Hexa", strconv.FormatInt(int64(ho.subject.getState()), 6))
}

func ObserveNumber() {
	s := Subject{state: 15}
	b := BinaryObserver{&s}
	b.subject.attach(b)
	h := HexaObserver{&s}
	h.subject.attach(h)
	s.setState(13)
}
