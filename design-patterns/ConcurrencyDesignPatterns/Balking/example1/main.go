package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	In diesem Beispiel habe ich eine SomeData-Struktur,
	welches ein Datenfeld und ein mu-Feld, das ein sync.Mutex
	ist, beinhaltet, um den Zugriff auf die Daten zu schützen.

	Die Methode changeData Methode aktualisiert die Daten, wenn sie
	nicht bereits auf den neuen Wert gesetzt sind, andernfalls bricht sie ab
	und kehrt vorzeitig zurück.

	In der Hauptfunktion erstelle ich eine SomeData-Instanz und starte fünf
	sog. Goroutinen, die changeData mit verschiedenen Datenwerten aufrufen,
	nachdem sie eine bestimmte Anzahl von Sekunden im Schlafmodus sind.
	Die time.Sleep-Aufrufe am Ende sowohl der Goroutine als auch der Hauptfunktion
	stellen sicher, dass alle Goroutinen abgeschlossen sind, bevor das Programm
	beendet wird.
*/

type SomeData struct {
	data string
	mu   sync.Mutex
}

func (sd *SomeData) changeData(newData string) {
	sd.mu.Lock()
	defer sd.mu.Unlock()

	if sd.data == newData {
		fmt.Println("Data is already set to", newData)
		return
	}

	fmt.Println("Setting data to", newData)
	sd.data = newData
}

func main() {
	sd := &SomeData{data: ""}

	for i := 0; i < 5; i++ {
		go func(i int) {
			newData := fmt.Sprintf("data%d", i)
			time.Sleep(time.Duration(i) * time.Second)
			sd.changeData(newData)
		}(i)
	}
	time.Sleep(6 * time.Second)
}
