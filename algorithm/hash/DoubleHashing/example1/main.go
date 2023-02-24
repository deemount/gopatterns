package main

import "log"

/*

	Beispiel:

	In dieser Implementierung definiere ich eine HashTable-Struktur mit einer Zuordnung
	von Integer-Schlüsseln zu Node-Zeigern. Ich definiere zudem eine Node-Struktur,
	die ein Schlüssel-Wert-Paar enthält. Die Hash-Tabelle wird mit der Funktion NewHashTable
	initialisiert, die einen Zeiger auf eine neu zugewiesene HashTable-Struktur zurückgibt.

	Die Hash-Tabelle verwendet zwei Hash-Funktionen, hash1 und hash2. Die erste Hash-Funktion
	ist eine einfache Modulo-Arithmetik-Operation, während die zweite Hash-Funktion eine komplexere
	Formel ist, die den Schlüsselwert und eine Konstante verwendet.

	Die Funktion put fügt ein neues Schlüssel-Wert-Paar in die Hash-Tabelle ein, wobei zur Auflösung
	von Kollisionen ein doppeltes Hashing verwendet wird. Die get-Funktion ruft den mit einem
	bestimmten Schlüssel verknüpften Wert ab, wiederum unter Verwendung von Double Hashing zur
	Suche nach dem Schlüssel. Die delete-Funktion entfernt ein Schlüssel-Wert-Paar aus der Hash-Tabelle und verwendet dabei Double Hashing, um den Schlüssel zu finden.

*/

// Define a hash table struct
type HashTable struct {
	table map[int]*Node
}

// Define a node struct
type Node struct {
	key   int
	value string
}

// Initialize the hash table
func NewHashTable() *HashTable {
	return &HashTable{make(map[int]*Node)}
}

// Define the first hash function
func (ht *HashTable) hash1(key int) int {
	return key % 10 // simple modulo arithmetic
}

// Define the second hash function
func (ht *HashTable) hash2(key int) int {
	return 7 - (key % 7) // second hash function
}

// Define the put function to insert a key-value pair
func (ht *HashTable) Put(key int, value string) {
	node := &Node{key, value}
	index := ht.hash1(key)
	if ht.table[index] == nil {
		ht.table[index] = node
	} else {
		i := 1
		for {
			newIndex := (index + i*ht.hash2(key)) % 10 // double hash function
			if ht.table[newIndex] == nil {
				ht.table[newIndex] = node
				break
			}
			i++
		}
	}
}

// Define the get function to retrieve a value given a key
func (ht *HashTable) Get(key int) string {
	index := ht.hash1(key)
	if ht.table[index] != nil && ht.table[index].key == key {
		return ht.table[index].value
	} else {
		i := 1
		for {
			newIndex := (index + i*ht.hash2(key)) % 10 // double hash function
			if ht.table[newIndex] != nil && ht.table[newIndex].key == key {
				return ht.table[newIndex].value
			} else if ht.table[newIndex] == nil {
				return "" // key not found
			}
			i++
		}
	}
}

// Define the delete function to remove a key-value pair
func (ht *HashTable) Delete(key int) {
	index := ht.hash1(key)
	if ht.table[index] != nil && ht.table[index].key == key {
		ht.table[index] = nil
	} else {
		i := 1
		for {
			newIndex := (index + i*ht.hash2(key)) % 10 // double hash function
			if ht.table[newIndex] != nil && ht.table[newIndex].key == key {
				ht.table[newIndex] = nil
				break
			} else if ht.table[newIndex] == nil {
				break // key not found
			}
			i++
		}
	}
}

func main() {
	ht := NewHashTable()

	// Insert some key-value pairs using the put function
	ht.Put(10, "Hello")
	ht.Put(23, "World")
	ht.Put(15, "Goodbye")

	// Retrieve values using the get function
	log.Println(ht.Get(10)) // Output: Hello
	log.Println(ht.Get(23)) // Output: World
	log.Println(ht.Get(15)) // Output: Goodbye
	log.Println(ht.Get(20)) // Output: (empty string, key not found)

	// Remove a key-value pair using the delete function
	ht.Delete(23)

	// Try to retrieve the value that was deleted
	log.Println(ht.Get(23)) // Output: (empty string, key not found)
}
