package main

import (
	"fmt"
)

/*

	Beispiel:

	In dieser Implementierung stellt der Typ hashTable die Cuckoo-Hashtabelle dar.
	Sie hat eine feste Größe, die in zwei Tabellen unterteilt ist.
	Die Hash-Methode berechnet den Hash-Wert einer Eingabe x auf der Grundlage der Tabellennummer.
	Die insert-Methode fügt ein neues Element x in die Tabelle ein und verwendet bis zu maxAttempts,
	um eventuelle Kollisionen durch das Entfernen vorhandener Elemente aufzulösen.
	Die search-Methode sucht nach einem Element x in der Tabelle.

	In der Hauptfunktion erstellen wir eine neue Hashtabelle und fügen einige Elemente ein.
	Anschließend wird nach einigen Elementen gesucht, um zu überprüfen, ob sie in der Tabelle vorhanden sind.

*/

const (
	maxAttempts = 10
)

type hashTable struct {
	table [2][]int
	size  int
}

func (h *hashTable) hash(x int, tableNum int) int {
	if tableNum == 0 {
		return x % (h.size / 2)
	}
	return (x / (h.size / 2)) % (h.size / 2)
}

func (h *hashTable) insert(x int) bool {
	if h.search(x) {
		return true
	}
	attempts := 0
	for attempts < maxAttempts {
		tableNum := attempts % 2
		pos := h.hash(x, tableNum)
		if h.table[tableNum][pos] == 0 {
			h.table[tableNum][pos] = x
			return true
		}
		evicted := h.table[tableNum][pos]
		h.table[tableNum][pos] = x
		x = evicted
		attempts++
	}
	return false
}

func (h *hashTable) search(x int) bool {
	for i := 0; i < 2; i++ {
		pos := h.hash(x, i)
		if h.table[i][pos] == x {
			return true
		}
	}
	return false
}

func main() {
	h := hashTable{
		table: [2][]int{
			make([]int, 5),
			make([]int, 5),
		},
		size: 10,
	}
	h.insert(10)
	h.insert(20)
	h.insert(30)
	h.insert(5)
	h.insert(15)
	h.insert(25)
	h.insert(35)
	h.insert(8)
	h.insert(18)
	h.insert(28)
	fmt.Println(h.search(10))  // true
	fmt.Println(h.search(15))  // true
	fmt.Println(h.search(25))  // true
	fmt.Println(h.search(35))  // true
	fmt.Println(h.search(100)) // false
}
