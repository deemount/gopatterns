package main

import "log"

/*

	Beispiel:

	In diesem Beispiel definiere ich eine SuccinctDict-Struktur, die ein Trie-basiertes
	succinct dictionary darstellt. Das Trie-Feld ist ein Slice von Slices von Ganzzahlen,
	wobei jedes Slice einen Knoten im Trie darstellt und Paare von (Rune, Kindindex) Werten enthält.
	Der Kindindex ist ein Index in der Trie-Scheibe, der den Index des Kindknotens darstellt.

	Ich definiere Insert- und Lookup-Methoden, die auf der Trie operieren, indem sie neue Schlüssel
	einfügen und indem sie einem Pfad durch die Trie folgen, Schlüssel nachschlagen und nach einem
	Kindknoten am Ende des Pfades suchen. Wir definieren auch einen NewSuccinctDict-Konstruktor,
	die ein neues, leeres SuccinctDict erzeugt.

	In der Hauptfunktion erstelle ich ein neues SuccinctDict und füge eine Reihe von Schlüsseln ein.
	Anschließend wird das Wörterbuch nach jedem der Schlüssel abgefragt, um zu zeigen, dass sie im
	Wörterbuch vorhanden sind.

	Diese Implementierung ist ein einfaches Beispiel für ein SuccinctDict, das noch weiter optimiert
	werden könnte, um die Leistung oder den Speicherverbrauch zu erhöhen. Sie sollte jedoch ausreichen,
	um die Kernkonzepte eines trie-basierten SuccinctDict zu demonstrieren.

*/

// SuccinctDict is a simple implementation of a succinct dictionary using a trie.
type SuccinctDict struct {
	trie [][]int
}

// NewSuccinctDict creates a new, empty SuccinctDict.
func NewSuccinctDict() *SuccinctDict {
	return &SuccinctDict{
		trie: [][]int{{}},
	}
}

// Insert adds a key to the dictionary.
func (d *SuccinctDict) Insert(key string) {
	node := 0
	for _, c := range key {
		child := -1
		for i := 0; i < len(d.trie[node]); i += 2 {
			if rune(d.trie[node][i]) == c {
				child = d.trie[node][i+1]
				break
			}
		}
		if child == -1 {
			child = len(d.trie)
			d.trie = append(d.trie, []int{})
			d.trie[node] = append(d.trie[node], int(c), child)
		}
		node = child
	}
	// Mark the end of the key with a final child node.
	d.trie = append(d.trie, []int{})
	d.trie[node] = append(d.trie[node], -1, len(d.trie)-1)
}

// Lookup returns the index of the key in the dictionary, or -1 if the key is not present.
func (d *SuccinctDict) Lookup(key string) int {
	node := 0
	for _, c := range key {
		child := -1
		for i := 0; i < len(d.trie[node]); i += 2 {
			if rune(d.trie[node][i]) == c {
				child = d.trie[node][i+1]
				break
			}
		}
		if child == -1 {
			return -1
		}
		node = child
	}
	// Check if the final node of the key is a final child node.
	for i := 0; i < len(d.trie[node]); i += 2 {
		if d.trie[node][i] == -1 {
			return d.trie[node][i+1]
		}
	}
	return -1
}

func main() {
	keys := []string{"apple", "banana", "cherry", "grape", "orange", "pear", "pineapple", "strawberry", "watermelon"}

	// Build the succinct dictionary from the keys
	dict := NewSuccinctDict()
	for _, key := range keys {
		dict.Insert(key)
	}

	// Query the dictionary for keys
	for _, key := range keys {
		if dict.Lookup(key) == -1 {
			log.Printf("Key '%s' not found in dictionary\n", key)
		} else {
			log.Printf("Key '%s' found in dictionary\n", key)
		}
	}
}
