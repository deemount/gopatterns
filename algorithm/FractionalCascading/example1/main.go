package main

import (
	"fmt"
)

/*

	Beispiel:

	In diesem Beispiel nimmt die Funktion FractionalCascading zwei sortierte Arrays A und B
	und führt eine fraktionierte Kaskadierung an ihnen durch, um alle Elemente zu finden,
	die in beiden Arrays vorkommen.

	Die Funktion erstellt zunächst eine verknüpfte Liste von Knoten für jedes Array,
	wobei jeder Knoten ein Element in dem Array darstellt.

	Das Feld Next jedes Knotens zeigt auf den nächsten Knoten in der verknüpften Liste mit
	demselben Wert. Die Funktion führt dann eine Zusammenführung der beiden verknüpften
	Listen durch und fügt dem Ergebnisfeld Elemente in sortierter Reihenfolge hinzu.

*/

// Node represents a node in the Fractional Cascading tree.
type Node struct {
	Value int
	Next  *Node
}

// FractionalCascading performs fractional cascading on two sorted arrays, A and B.
// It returns an array containing all the elements that appear in both A and B.
func FractionalCascading(A, B []int) []int {
	// Create a list of nodes for each array, including a sentinel node at the end.
	nodesA := make([]*Node, len(A)+1)
	nodesB := make([]*Node, len(B)+1)
	for i := range nodesA {
		nodesA[i] = &Node{}
	}
	for i := range nodesB {
		nodesB[i] = &Node{}
	}

	// Fill in the nodes with the corresponding values.
	for i := len(A) - 1; i >= 0; i-- {
		nodesA[i].Value = A[i]
		if i == len(A)-1 {
			nodesA[i].Next = nodesA[i+1]
		} else if A[i] == A[i+1] {
			nodesA[i].Next = nodesA[i+1]
		} else {
			nodesA[i].Next = nodesA[i+1]
		}
	}
	for i := len(B) - 1; i >= 0; i-- {
		nodesB[i].Value = B[i]
		if i == len(B)-1 {
			nodesB[i].Next = nodesB[i+1]
		} else if B[i] == B[i+1] {
			nodesB[i].Next = nodesB[i+1]
		} else {
			nodesB[i].Next = nodesB[i+1]
		}
	}

	// Perform the merge.
	result := []int{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if nodesA[i].Value == nodesB[j].Value {
			result = append(result, nodesA[i].Value)
			i++
			j++
		} else if nodesA[i].Value < nodesB[j].Value {
			result = append(result, nodesA[i].Value)
			i++
		} else {
			result = append(result, nodesB[j].Value)
			j++
		}
	}

	// Add any remaining elements from A or B.
	for i < len(A) {
		result = append(result, nodesA[i].Value)
		i++
	}
	for j < len(B) {
		result = append(result, nodesB[j].Value)
		j++
	}

	return result
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := []int{3, 4, 5, 6, 7}
	result := FractionalCascading(A, B)
	fmt.Println(result) // Ausgabe: [3 4 5]
}
