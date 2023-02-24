package main

/*

	Beispiel:

	In diesem Beispiel suche ich nach dem Zielwert 5 in der sortierten Matrix arr.
	Die Funktion gibt den Index des Zielwerts zurück, wenn er gefunden wurde, oder -1,
	wenn er sich nicht im Array befindet. Die Implementierung basiert auf der Idee, dass,
	wenn die Elemente des Arrays gleichmäßig verteilt sind, ich die Position des Zielwerts
	auf der Grundlage seines Werts relativ zu den Minimal- und Maximalwerten im Array schätzen
	kann.

*/

import "log"

func InterpolationSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right && target >= arr[left] && target <= arr[right] {
		pos := left + ((target-arr[left])*(right-left))/(arr[right]-arr[left])
		if arr[pos] == target {
			return pos
		}
		if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	target := 14
	index := InterpolationSearch(arr, target)
	log.Println(index) // Ausgabe: 13

}
