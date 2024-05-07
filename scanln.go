package scanln

import (
	"fmt"
	"os"
)

type Scanln struct {
	C chan string
	q chan bool
}

func NewScanln() *Scanln {
	s := &Scanln{make(chan string), make(chan bool, 1)}
	go s.runForever()

	return s
}

func (s *Scanln) Stop() {
	s.q <- true

	file := os.NewFile(0, "stdin")
	_, err := file.Write([]byte(``))
	if err != nil {
		panic(err)
	}
}

func (s *Scanln) runOne() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func (s *Scanln) runForever() {
	for {
		input := s.runOne()
		select {
		case <-s.q:
			return
		default:
			s.C <- input
		}
	}
}
