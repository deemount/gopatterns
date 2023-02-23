package main

/*

	Ich erstelle einen Kanal c, der zur Übermittlung des Ergebnisses einer langlaufenden Aufgabe verwendet werden soll.
	Anschließend erstelle ich eine Goroutine, die die langlaufende Aufgabe simuliert und das Ergebnis an den Kanal sendet,
	wenn sie fertig ist.

	Als Nächstes verwende ich eine select-Anweisung, um entweder auf das Ergebnis des Kanals oder auf eine Zeitüberschreitung
	von einer Sekunde zu warten. Wenn das Ergebnis vor der Zeitüberschreitung eintrifft, wird es auf der Konsole ausgegeben.
	Andernfalls wird eine Timeout-Meldung ausgegeben.

	Durch die Verwendung eines Kanals zur Übermittlung des Ergebnisses und einer Select-Anweisung zum Warten darauf haben wir
	das bewachte Guarded-Suspension Entwurfsmuster nun implementiert. Die Haupt-Goroutine wartet, bis das Ergebnis verfügbar ist,
	bevor sie mit ihrer Aufgabe fortfährt, die in diesem Fall darin besteht, das Ergebnis auf der Konsole auszugeben.

	Beachte, dass der in der Funktion time.After verwendete Timeout-Wert an die Anforderungen des jeweiligen Anwendungsfalls
	angepasst werden kann.

*/

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		// simulate a long-running task
		time.Sleep(2 * time.Second)
		c <- "result"
	}()

	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
