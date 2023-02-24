package main

import (
	"log"
	"sort"
)

/*
	kth smallest number

	Beispiel:

	In diesem Algorithmus nehme ich ein ganzzahliges Array arr und eine ganze Zahl k als Eingabe.
	Ich sortiere das Array zunächst mit der Funktion sort.Ints, die von Golangs eingebautem sort-Paket
	bereitgestellt wird. Dann gebe ich das k-te Element des sortierten Arrays zurück, indem wir auf
	arr[k-1] zugreifen.

	Beachte, dass dieser Algorithmus annimmt, dass k ein gültiger Index im Array ist, was bedeutet,
	dass k größer oder gleich 1 und kleiner oder gleich der Länge des Arrays sein muss.
	Wenn k außerhalb des Bereichs liegt, bricht der Algorithmus mit einem Index außerhalb des
	Bereichs-Fehler ab. Gegebenenfalls sollte der Algorithmus eine Fehlerbehandlung hinzufügen,
	um diesen Fall zu behandeln.

*/

func kthSmallest(arr []int, k int) int {
	// Sort the array in ascending order
	sort.Ints(arr)

	// Return the kth element
	return arr[k-1]
}

func main() {
	result := kthSmallest([]int{23, 65, 42, 2, 5, 90, 100, 34, 3, 6, 7, 8, 112}, 5)
	log.Printf("main: result %d", result)
}
