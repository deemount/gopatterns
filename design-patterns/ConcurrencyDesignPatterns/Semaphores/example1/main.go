package main

import (
	"errors"
	"time"
)

/*

	Beispiel:

	Beschreibung folgt ...

*/

var (
	ErrNoTickets      = errors.New("semaphore: kann keine semaphore erwerben")
	ErrIllegalRelease = errors.New("semaphore: kann semaphore nicht ausgeben bevor sie nicht erweorben worden ist")
)

// Die Schnittstelle enthält das Verhalten einer Semaphore,
// die erworben und/oder freigegeben werden kann.
type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

func main() {

	// Semaphore with Timeouts

	tickets, timeout := 25, 10*time.Second
	s := New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	// hier kann weiterer code ausgeführt werden

	if err := s.Release(); err != nil {
		panic(err)
	}

	// Semaphore without Timeouts (Non-Blocking)
	/*
		tickets2, timeout2 := 0, 0*time.Second
		s2 := New(tickets2, timeout2)

		if err := s2.Acquire(); err != nil {
			if err != ErrNoTickets {
				panic(err)
			}

			// No tickets left, can't work :(
			os.Exit(1)
		}
	*/

}
