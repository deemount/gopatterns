package main

import "log"

/*

	Beispiel:

	In diesem Beispiel definiere ich eine Funktion mit dem Namen linearSearch,
	die ein Stück Ganzzahl (arr) und eine Ziel-Ganzzahl (target) als Argumente annimmt.
	Die Funktion durchläuft den Slice mit einer for-Schleife und prüft, ob jedes Element
	des Slice gleich der Ziel-Ganzzahl ist.

	* Wenn die Ziel-Ganzzahl gefunden wird, gibt die Funktion den Index zurück, bei dem sie gefunden wurde.
	* Wenn die Ziel-Ganzzahl nicht gefunden wird, gibt die Funktion -1 zurück.

	In der Hauptfunktion erstelle ich einen Slice von Ganzzahlen namens numbers und eine
	Ziel-Ganzzahl namens target. Dann rufe ich die Funktion linearSearch mit diesen Argumenten
	auf und speicher das Ergebnis in einer Variablen namens index.

	Schließlich überprüfe ich den Wert von index, um festzustellen, ob die Ziel-Ganzzahl im
	Slice gefunden wurde oder nicht und gebe eine entsprechende Meldung auf der Konsole aus.

*/

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{5, 10, 15, 20, 25}
	target := 15

	index := linearSearch(numbers, target)

	if index == -1 {
		log.Printf("%d was not found in the slice\n", target)
	} else {
		log.Printf("%d was found at index %d in the slice\n", target, index)
	}
}
