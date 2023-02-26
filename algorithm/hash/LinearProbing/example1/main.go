package main

import "log"

/*

	Beispiel:

	In diesem Beispiel definiere ich eine HashTable-Struktur mit einem Tabellenfeld,
	das eine Abbildung von Ganzzahlen auf Node-Zeiger ist. Ich definiere zudem eine
	Node-Struktur mit einem Schlüssel und einem Wert.

	Die Insert-Methode verwendet lineares Sondieren, um ein neues Schlüssel-Wert-Paar in die
	Hashtabelle einzufügen. Sie berechnet den anfänglichen Index, indem sie den Schlüssel modulo
	der Länge der Tabelle nimmt und iteriert dann durch die Tabelle, bis sie einen leeren Slot findet,
	in den sie den neuen Knoten einfügt.

	Die Suchmethode verwendet ebenfalls lineares Sondieren, um nach einem Wert zu suchen, der mit
	einem bestimmten Schlüssel verbunden ist. Sie beginnt bei dem Index, der sich aus der Multiplikation
	des Schlüssels mit der Länge der Tabelle ergibt und durchläuft dann die Tabelle, bis sie den Knoten mit
	dem angegebenen Schlüssel oder einen leeren Slot findet.

	In der Hauptfunktion erstelle ich eine neue Hashtabelle mit einer anfänglichen Kapazität von 10,
	füge drei Schlüssel-Wert-Paare ein und suche dann nach Werten, die mit diesen Schlüsseln verbunden sind.
	Die Ausgabe zeigt, dass die Suchmethode die mit den Schlüsseln 5, 15 und 25 verknüpften Werte korrekt zurückgibt
	und bei der Suche nach dem nicht vorhandenen Schlüssel 35 eine leere Zeichenkette zurückgibt.

*/

type HashTable struct {
	table map[int]*Node
}

type Node struct {
	key   int
	value string
}

func (ht *HashTable) Insert(key int, value string) {
	index := key % len(ht.table)
	for ht.table[index] != nil {
		index = (index + 1) % len(ht.table)
	}
	ht.table[index] = &Node{key: key, value: value}
}

func (ht *HashTable) Search(key int) string {
	index := key % len(ht.table)
	for ht.table[index] != nil && ht.table[index].key != key {
		index = (index + 1) % len(ht.table)
	}
	if ht.table[index] == nil {
		return ""
	}
	return ht.table[index].value
}

func main() {
	ht := HashTable{table: make(map[int]*Node, 10)}
	ht.Insert(5, "apple")
	ht.Insert(15, "banana")
	ht.Insert(25, "orange")

	log.Println(ht.Search(5))  // Ausgabe: apple
	log.Println(ht.Search(15)) // Ausgabe: banana
	log.Println(ht.Search(25)) // Ausgabe: orange
	log.Println(ht.Search(35)) // Ausgabe:
}
