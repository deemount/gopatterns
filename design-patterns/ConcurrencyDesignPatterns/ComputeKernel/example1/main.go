package main

/*

	Beispiel:

	Ich erstelle zwei Kanäle: Aufträge und Ergebnisse.
	Ich erzeuge eine zufällige Arbeit, indem ich 100 Aufträge an den Auftragskanal sende
	und ihn dann schließe, um zu signalisieren, dass keine weiteren Aufträge mehr gesendet
	werden sollen.

	Dann starte ich vier Worker mit Hilfe von Goroutinen, die jeweils Aufträge aus dem Auftragskanal
	lesen, einige umfangreiche Berechnungen mit dem Auftrag durchführen und dann das Ergebnis über den
	Ergebniskanal an den Hauptthread zurücksenden.

	Schließlich erstelle ich eine separate Goroutine, um die Ergebnisse aus dem Ergebniskanal zu sammeln,
	wobei ich eine sync.WaitGroup verwende, um sicherzustellen, dass alle Worker den Auftrag abgeschlossen haben,
	bevor wir den Ergebniskanal schließen und das Programm beenden.

	Diese Implementierung des Rechenkernmusters in Go ermöglicht es uns, die umfangreichen Berechnungen auf
	mehrere Worker zu verteilen, was die Leistung verbessert und die Gesamtberechnungszeit reduziert.

*/

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	numWorkers = 4
	numJobs    = 100
)

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// generate some random work
	for i := 0; i < numJobs; i++ {
		jobs <- rand.Intn(100)
	}
	close(jobs)

	// start the workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// collect the results
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// perform some heavy computation on the job
		result := job * 2
		// send the result back to the main thread
		results <- result
	}
}
