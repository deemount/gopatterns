package main

import "log"

/*

	Beispiel:

	In dieser Implementierung ist MinHeap ein Slice von Ganzzahlen.
	Die Push-Methode fügt dem Heap ein neues Element hinzu und ordnet
	dann die Elemente neu an, um die MinHeap-Eigenschaft zu erhalten,
	die besagt, dass das übergeordnete Element jedes Knotens kleiner
	oder gleich seinen Kindern sein muss.

	Die Pop-Methode entfernt das kleinste Element aus dem Heap und gibt
	es zurück. Anschließend ordnet sie die verbleibenden Elemente neu an,
	um die Eigenschaft min heap zu wahren.

	Die Methode minHeapify ist eine Hilfsfunktion, die die Elemente neu anordnet,
	um die Eigenschaft min heap beizubehalten, beginnend bei einem gegebenen Index i
	und aufsteigend bis zum Index n. Sie tut dies, indem sie die Elemente nach Bedarf
	vertauscht, so dass das kleinste Element immer das übergeordnete Element seiner
	Kinder ist.

*/

type MinHeap []int

func (h *MinHeap) Push(x int) {
	*h = append(*h, x)
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		h.minHeapify(i, n)
	}
}

func (h *MinHeap) Pop() int {
	n := len(*h)
	x := (*h)[0]
	(*h)[0], (*h)[n-1] = (*h)[n-1], (*h)[0]
	*h = (*h)[:n-1]
	h.minHeapify(0, n-1)
	return x
}

func (h *MinHeap) minHeapify(i, n int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && (*h)[left] < (*h)[smallest] {
			smallest = left
		}
		if right < n && (*h)[right] < (*h)[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}

		(*h)[i], (*h)[smallest] = (*h)[smallest], (*h)[i]
		i = smallest
	}
}

// main
func main() {

	var numbers MinHeap = []int{12, 34, 89, 1, 4, 23, 90, 87, 66, 45, 2}
	log.Printf("main: numbers %v", numbers)
	numbers.Push(101)
	log.Printf("main: numbers (after push) %v", numbers)
	numbers.Pop()
	log.Printf("main: numbers (after pop) %v", numbers)
	numbers.Pop()
	numbers.Pop()
	numbers.Pop()
	numbers.Pop()
	log.Printf("main: numbers (after pop) %v", numbers)
	numbers.Push(101)
	numbers.Push(102)
	numbers.Push(103)
	numbers.Push(104)
	log.Printf("main: numbers (after push) %v", numbers)
}
