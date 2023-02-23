package main

import (
	"fmt"
	"sync"
)

/*

	Beispiel:

	Ich definiere eine SafeCounter-Struktur, die ein sync.Mutex-Feld enthält,
	das zur Synchronisierung des Zugriffs auf das Zählfeld verwendet wird.
	Die Methode Increment() erwirbt die Sperre, erhöht den Zählerstand und gibt
	die Sperre wieder frei. Die Methode GetCount() übernimmt ebenfalls die Sperre,
	gibt den Zählerstand zurück und gibt dann die Sperre mit einer Defer-Anweisung wieder frei,
	um sicherzustellen, dass die Sperre immer freigegeben wird, auch wenn es zu einer Panik kommt.

	Anschließend wird eine SafeCounter-Instanz erstellt und 1000 Goroutinen gestartet, von denen
	jede den Zähler durch Aufruf der Increment()-Methode inkrementiert. Zum Schluss warten wir, bis alle
	Goroutinen beendet sind, und geben den endgültigen Zählerstand aus.

	Diese Implementierung stellt sicher, dass jeweils nur eine Goroutine auf das Zählfeld zugreifen kann,
	wodurch Wettlaufsituationen verhindert und eine korrekte Synchronisierung des Zugriffs auf die gemeinsame
	Ressource gewährleistet wird.

*/

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (s *SafeCounter) Increment() {
	s.mu.Lock()
	s.count++
	s.mu.Unlock()
}

func (s *SafeCounter) GetCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

func main() {
	counter := SafeCounter{count: 0}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	fmt.Println(counter.GetCount())
}
