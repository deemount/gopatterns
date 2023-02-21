package main

/*
	In diesem Beispiel wird das ereignisbasierte asynchrone Muster verwendet,
	um Ereignisse auf nicht blockierende Weise zu verarbeiten. Das Programm
	beginnt mit der Erstellung von zwei Kanälen, eventChan und doneChan.
	Der eventChan-Kanal wird zum Senden von Ereignissen verwendet und der
	doneChan-Kanal wird verwendet, um zu signalisieren, dass alle Ereignisse
	gesendet wurden.

	Mit dem Schlüsselwort go werden zwei Goroutinen erstellt, die gleichzeitig
	ablaufen. Die erste Goroutine ist der Producer, der jede Sekunde Ereignisse
	an den eventChan-Kanal sendet. Die zweite Goroutine ist der Consumer, der
	Ereignisse aus dem eventChan-Kanal empfängt und verarbeitet.

	Die select-Anweisung wird verwendet, um Ereignisse aus dem eventChan-Kanal
	zu empfangen und ein Signal aus dem doneChan-Kanal zu erhalten. Auf diese
	Weise kann das Programm Ereignisse verarbeiten, sobald sie empfangen werden,
	ohne den Kanal zu blockieren. Sobald der doneChan-Kanal ein Signal empfängt,
	wird die verbraucherende Goroutine beendet.

	Schließlich wartet das Programm sechs Sekunden lang, damit alle Ereignisse gesendet
	und empfangen werden können, bevor es beendet wird.
*/

import (
	"fmt"
	"time"
)

type event struct {
	id    int
	value string
}

func main() {

	eventChan := make(chan event, 10)
	doneChan := make(chan bool)

	// producer
	go func() {
		for i := 1; i <= 5; i++ {
			eventChan <- event{id: i, value: fmt.Sprintf("Event %d", i)}
			time.Sleep(time.Second)
		}
		doneChan <- true
	}()

	// consumer
	go func() {
		for {
			select {
			case evt := <-eventChan:
				fmt.Printf("Event Received: %v\n", evt)
			case <-doneChan:
				fmt.Println("Done receiving events")
				return
			}
		}
	}()

	fmt.Println("Waiting for events to be received...")
	time.Sleep(6 * time.Second)
}
