package main

import "fmt"

/*

	Beispiel:

	In diesem Beispiel habe ich zwei Funktionen: exponentialSearch und binarySearch.

	exponentialSearch führt den exponentiellen Suchalgorithmus durch, indem es zunächst prüft,
	ob sich das Zielelement an der ersten Position des Arrays befindet. Ist dies der Fall,
	gibt die Funktion 0 zurück. Andernfalls setzt sie einen Anfangsindex i auf 1 und verdoppelt ihn,
	bis i entweder größer oder gleich der Länge des Arrays ist oder das Element bei Index i größer als
	das Zielelement ist.

	Dann ruft es binarySearch auf, um eine binäre Suche auf dem Subarray zwischen den Indizes i/2 und
	min(i, n-1) durchzuführen.

	binarySearch ist eine Hilfsfunktion, die eine binäre Suche auf einem gegebenen Unterarray des Eingabe-Arrays
	durchführt. Sie nimmt vier Parameter entgegen: das Eingabe-Array arr, die linken und rechten Indizes des
	Subarrays und das Zielelement x. Sie beginnt mit der Berechnung des mittleren Index des Subarrays und
	ergleicht das Element bei diesem Index mit dem Zielelement.

	Wenn das mittlere Element gleich dem Zielelement ist, gibt die Funktion den mittleren Index zurück.
	Andernfalls verkleinert sie den Teilbereich, indem sie entweder den linken Index auf den mittleren
	Index + 1 oder den rechten Index auf den mittleren Index - 1 verschiebt, je nachdem, ob das mittlere
	Element kleiner oder größer als das Zielelement ist.

	Dieser Vorgang wird so lange wiederholt, bis das Zielelement gefunden wurde oder das Subarray leer ist.

	In der Hauptfunktion definiere ich ein Beispielarray arr, ein Zielelement x und rufe exponentialSearch auf,
	um den Index des Zielelements im Array zu finden.

	Wenn der Index nicht -1 ist, gebe ich eine Meldung aus, dass das Zielelement bei diesem Index gefunden wurde.
	Andernfalls gebe ich eine Meldung aus, dass das Zielelement nicht im Array gefunden wurde.

*/

func exponentialSearch(arr []int, x int) int {
	n := len(arr)

	if arr[0] == x {
		return 0
	}

	i := 1
	for i < n && arr[i] <= x {
		i = i * 2
	}

	return binarySearch(arr, i/2, min(i, n-1), x)
}

func binarySearch(arr []int, left int, right int, x int) int {
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{2, 5, 7, 10, 13, 16, 18}
	x := 13
	index := exponentialSearch(arr, x)
	if index != -1 {
		fmt.Printf("Found %d at index %d\n", x, index)
	} else {
		fmt.Printf("%d not found in the array\n", x)
	}
}
