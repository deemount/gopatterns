package main

/*

	Beispiel:

	In diesem Beispiel ist die Zählervariable eine gemeinsam genutzte Ressource,
	die vor gleichzeitigem Zugriff geschützt werden muss. Zu diesem Zweck erstelle
	ich eine sync.Mutex-Variable namens mutex.

	In der Inkrementfunktion wird die Sperre durch den Aufruf von mutex.Lock() aktiviert,
	wodurch andere Goroutinen am Zugriff auf die gemeinsame Ressource gehindert werden.
	Sobald ich fertig bin, gebe ich die Sperre durch den Aufruf von mutex.Unlock() wieder
	frei. Ich verwende das Schlüsselwort defer, um sicherzustellen, dass der Mutex immer
	entsperrt ist, auch wenn während der Ausführung ein Fehler auftritt.

	Um den Mutex zu testen, erzeuge ich 100 Goroutinen, die die Inkrement-Funktion aufrufen
	und verwende eine sync.WaitGroup, um zu warten, bis alle von ihnen fertig sind. Der Endwert
	des Zählers sollte 100 sein, was beweist, dass der Mutex die gemeinsame Ressource erfolgreich
	vor dem gleichzeitigen Zugriff geschützt hat.

*/

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	wg.Done()
}
