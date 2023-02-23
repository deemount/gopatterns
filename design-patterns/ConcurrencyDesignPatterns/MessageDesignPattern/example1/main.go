package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

	Beispiel:

	Ich erstelle ein einfaches Nachrichtensystem, in dem zwei Goroutinen über eine
	Nachrichtenwarteschlange miteinander kommunizieren. Die erste Goroutine generiert
	Zufallszahlen und sendet sie an die Nachrichtenwarteschlange, während die zweite
	Goroutine die Nachrichten aus der Warteschlange empfängt und sie auf der Konsole ausgibt.

	Ich definiere eine Message-Struktur, um die Nachrichten darzustellen, die zwischen den
	beiden Goroutinen gesendet werden sollen. Anschließend erstelle ich eine Nachrichtenwarteschlange
	mit der Funktion make, um einen Kanal vom Typ Message zu erzeugen.

	Ich starte zwei Goroutinen mit anonymen Funktionen. Die erste Goroutine erzeugt eine Zufallszahl
	zwischen 1 und 100 und erstellt eine Message-Struktur mit der Zufallszahl. Anschließend sendet sie
	die Nachricht mit dem Operator <- an die Nachrichtenwarteschlange. Ich verwende auch die Funktion
	time.Sleep, um eine zufällige Zeitspanne abzuwarten, bevor die nächste Nachricht gesendet wird.

	Die zweite Goroutine empfängt mit dem Operator <- Nachrichten aus der Nachrichtenwarteschlange und
	gibt den Wert der Nachricht mit der Funktion fmt.Println auf der Konsole aus.

	Schließlich verwende ich die Anweisung select {}, um die Haupt-Goroutine zu blockieren und um das
	Programm auf unbestimmte Zeit weiterlaufen zu lassen. Dies ist notwendig, da beide Goroutinen im
	Hintergrund laufen. Wenn die Hauptgoroutine beendet wird, wird das Programm beendet und die anderen
	Goroutinen werden nicht mehr ausgeführt.

*/

type Message struct {
	Value int
}

func main() {

	// create a channel for sending and receiving messages
	msgQueue := make(chan Message)

	// start a goroutine to generate messages
	go func() {
		for {
			// generate a random number between 1 and 100
			value := rand.Intn(100) + 1

			// create a message with the random number
			msg := Message{Value: value}

			// send the message to the message queue
			msgQueue <- msg

			// wait for a random amount of time before sending the next message
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()

	// start a goroutine to receive messages and print them to the console
	go func() {
		for {
			// receive a message from the message queue
			msg := <-msgQueue

			// print the value of the message to the console
			fmt.Println(msg.Value)
		}
	}()

	// block the main goroutine to keep the program running
	select {}

}
