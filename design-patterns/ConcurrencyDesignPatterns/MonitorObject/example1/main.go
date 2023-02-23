package main

import (
	"log"
	"sync"
)

/*

	Beispiel:

	Ich definiere eine Counter-Struktur, die einen Zähler, einen Mutex zum Schutz des Zählers
	und ein sync.Cond-Objekt enthält, um wartenden Goroutines zu signalisieren, wenn der Zähler
	erhöht wird.

	Die Increment-Methode erhöht den Zählerstand und signalisiert wartende Goroutinen mit der Methode
	cond.Signal().

	Die Methode WaitUntilGreaterThan wartet, bis der Zählerstand größer als ein bestimmter Wert ist.
	Sie verwendet eine Schleife, um den Zählerstand zu überprüfen und ruft cond.Wait() auf, um die Goroutine
	zu blockieren, bis der Zählerstand erhöht und die Bedingung erfüllt ist.

	Beachte, dass die Methode WaitUntilGreaterThan aufgerufen werden muss, während der Mutex gehalten wird,
	um eine Race Condition zu vermeiden.

*/

type Counter struct {
	count int
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewCounter() *Counter {
	c := &Counter{cond: sync.NewCond(&sync.Mutex{})}
	return c
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
	log.Printf("increment: count %d", c.count)
	c.cond.Signal()
}

func (c *Counter) WaitUntilGreaterThan(value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for c.count <= value {
		c.cond.Wait()
	}
}

func main() {

	counter := NewCounter()
	counter.count = 0
	for i := 0; i < 100; i++ {
		counter.Increment()
	}
	counter.Increment()
	counter.WaitUntilGreaterThan(100)
	for i := counter.count; i < 200; i++ {
		counter.Increment()
	}

}
