package main

import (
	"log"
	"math/rand"
	"time"
)

/*

	Beispiel:

	Diese Implementierung der binären Suche ist "verrauscht", weil sie Rauschen in der Eingabe simuliert,
	indem sie das Vorzeichen der Differenz zwischen dem Mittelpunkt und dem Ziel zufällig umkehrt. Dies wird erreicht,
	indem mit der Funktion rand.Intn() eine zufällige Ganzzahl zwischen 0 und 1 erzeugt und die Differenz mit -1
	multipliziert wird, wenn die zufällige Ganzzahl 0 ist.

	Beachte, dass diese Implementierung davon ausgeht, dass das Eingabefeld in aufsteigender Reihenfolge sortiert ist.
	Die Funktion gibt den Index des Zielelements zurück, wenn es gefunden wurde, oder -1, wenn es im Array nicht vorhanden ist.

	Beachte, dass das genaue Ergebnis aufgrund der zufälligen Natur des Rauschens abweichen kann.
	Manchmal führt das Rauschen dazu, dass die Suche über das Ziel hinausschießt, und manchmal dazu,
	dass sie unter das Ziel hinausschießt. In den meisten Fällen sollte der Algorithmus jedoch immer noch
	auf die richtige Antwort konvergieren.

*/

func NoisyBinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			// Simulating noise by randomly flipping the sign of the difference
			noisyDiff := target - arr[mid]
			if rand.Intn(2) == 0 {
				noisyDiff *= -1
			}
			low = mid + noisyDiff
		} else {
			// Simulating noise by randomly flipping the sign of the difference
			noisyDiff := arr[mid] - target
			if rand.Intn(2) == 0 {
				noisyDiff *= -1
			}
			high = mid - noisyDiff
		}
	}
	return -1
}

func main() {

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	arr := []int{1, 3, 4, 6, 8, 9, 11, 13, 15}
	target := 9

	log.Printf("Searching for %d in %v...\n", target, arr)

	index := NoisyBinarySearch(arr, target)

	if index == -1 {
		log.Printf("%d not found in %v\n", target, arr)
	} else {
		log.Printf("%d found at index %d in %v\n", target, index, arr)
	}

}
