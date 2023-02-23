package main

import (
	"fmt"
	"sync"
)

/*

	Beispiel:

	Ich erstellen eine WaitGroup-Variable namens wg. Dann durchlaufe ich fünf Schleifen
	und rufe die Add-Methode der WaitGroup auf, um den Zähler um 1 zu erhöhen.
	Ich rufe auch das go-Schlüsselwort auf, um eine neue Goroutine zu erstellen und
	übergebe ihr die Worker-Funktion und einen Verweis auf die WaitGroup.

	Innerhalb der Worker-Funktion verwende ich das Schlüsselwort defer, um die Done-Methode
	der WaitGroup aufzurufen, wenn die Funktion abgeschlossen ist. Dadurch wird der Zähler
	der WaitGroup dekrementiert.

	Schließlich rufen wir die Wait-Methode der WaitGroup für die Blockierung auf, bis alle
	Worker-Funktionen abgeschlossen sind. Sobald alle Worker fertig sind, drucken wir
	"All workers done" aus.

*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Do some work
	// ...

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers done")
}
