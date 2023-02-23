package main

import (
	"fmt"
	"sync"
)

/*

	Beispiel:

	Ich erstelle einen Kanal vom Typ struct{} mit einer Puffergröße von 3,
	um eine Semaphore mit einem Limit von 3 gleichzeitigen Arbeitern zu implementieren.
	Dann verwende ich eine sync.WaitGroup, um zu warten, bis alle Worker fertig sind,
	bevor ich das Programm beenden.

	In der for-Schleife werden 10 Worker-Goroutinen gestartet, und jede Goroutine erhält
	zunächst den Semaphor, indem sie ein leeres struct{} an den sem-Kanal sendet.
	Die Anweisung defer stellt sicher, dass die Semaphore freigegeben wird (durch Empfang vom sem-Kanal),
	wenn die Worker-Funktion beendet wird, unabhängig davon, ob sie normal oder aufgrund
	eines Fehlers beendet wird.

	Innerhalb der Worker-Funktion simulieren wir etwas Arbeit, indem wir eine Nachricht
	ausgeben, die anzeigt, dass die Worker-Funktion gestartet und beendet wurde, aber
	in einem realen Szenario würde dies durch die tatsächlich zu erledigende Arbeit
	ersetzt werden.

	Mit diesem Semaphor-Muster können bis zu 3 Worker gleichzeitig ausgeführt werden,
	aber alle weiteren Worker müssen warten, bis einer der aktuell ausgeführten Worker
	fertig ist und den Semaphor freigibt.

*/

func main() {
	sem := make(chan struct{}, 3) // erstelle einen kanal mit der größe 3
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			sem <- struct{}{} // semaphore erwerben
			defer func() {
				<-sem // semaphore ausgeben
				wg.Done()
			}()
			fmt.Printf("Worker %d gestarte\n", id)
			// hier kann weiterer code ausgeführt werden ...
			fmt.Printf("Worker %d beendet\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Alle workers beendet")
}
