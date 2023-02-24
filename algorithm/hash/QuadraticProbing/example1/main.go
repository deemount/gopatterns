package main

import (
	"fmt"
	"log"
)

/*

	Beispiel:

	In diesem Beispiel initialisiere ich zunächst eine Hashtabelle der Größe 7
	mit einem Slice in Golang. Dann füge einige Werte mit der Funktion insert ein,
	die eine quadratische Prüfung verwendet, um einen leeren Platz in der Hashtabelle
	zu finden. Wenn der anfängliche Hash-Index belegt ist, verwendet die Funktion eine
	quadratische Funktion, um den nächsten zu prüfenden Index zu berechnen.

	Ist auch dieser Platz belegt, verwendet die Funktion weiterhin die quadratische
	Funktion mit einem inkrementellen Wert, bis ein leerer Platz gefunden wird.

*/

func insert(table []int, value int) {
	// Calculate the initial hash index
	index := value % len(table)

	// If the slot is empty, insert the value
	if table[index] == 0 {
		table[index] = value
		return
	}

	// If the slot is not empty, use quadratic probing to find an empty slot
	i := 1
	for {
		// Calculate the next index using quadratic probing
		nextIndex := (index + i*i) % len(table)

		// If the next slot is empty, insert the value
		if table[nextIndex] == 0 {
			table[nextIndex] = value
			return
		}

		// If we've already checked every slot, the table is full
		if i == len(table) {
			fmt.Println("Table is full")
			return
		}

		// Increment the probing index
		i++
	}
}

func main() {
	// Initialize hash table with a size of 7
	table := make([]int, 7)

	// Insert some elements
	insert(table, 14)
	insert(table, 21)
	insert(table, 33)
	insert(table, 10)

	// Print the hash table
	log.Println(table)
}
