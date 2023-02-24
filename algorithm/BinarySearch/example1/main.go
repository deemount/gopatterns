package main

import "fmt"

/*

	Beispiel:

	In diesem Beispiel nimmt die Funktion binarySearch zwei Argumente entgegen:
	das sortierte Array arr und das Zielelement target. Die Funktion gibt den Index des
	 Zielelements im Array zurück, oder -1, wenn das Zielelement nicht gefunden wird.

	Der binäre Suchalgorithmus arbeitet, indem er das Suchintervall wiederholt in zwei
	Hälften teilt, bis das Zielelement gefunden ist. In jeder Iteration vergleicht der
	Algorithmus das mittlere Element des aktuellen Intervalls mit dem Zielelement.

	* Wenn das mittlere Element gleich dem Zielelement ist,
	gibt der Algorithmus den Index des mittleren Elements zurück.

	* Wenn das mittlere Element kleiner als das Zielelement ist,
	setzt der Algorithmus die Suche in der rechten Hälfte des aktuellen Intervalls fort.

	* Wenn das mittlere Element größer als das Zielelement ist,
	setzt der Algorithmus die Suche in der linken Hälfte des aktuellen Intervalls fort.

	In der Hauptfunktion erstelle ich ein sortiertes Array arr und ein Zielelement target mit dem Wert 7.
	Dann rufe ich die Funktion binarySearch mit diesen Argumenten auf und speicher das Ergebnis in der
	Variablen result. Zum Schluss geben wir eine Meldung aus, die angibt, ob das Zielelement gefunden wurde oder nicht.

*/

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Target element %d found at index %d", target, result)
	} else {
		fmt.Printf("Target element %d not found in the array", target)
	}
}
