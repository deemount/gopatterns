package main

import (
	"fmt"
)

/*

	Beispiel:

	In diesem Beispiel definiere ich eine Node-Struktur, die einen Schlüssel,
	einen Wert und einen Verweis auf den nächsten Knoten in der verknüpften Liste enthält.
	Ich definiere zudem eine HashTable-Struktur, die eine Größe und ein Array von
	Knotenzeigern namens Table enthält.

	Die Funktion NewHashTable erstellt eine neue Hash-Tabelle mit der angegebenen Größe
	und initialisiert das Table-Array mit Nullwerten.

	Die Hash-Funktion berechnet den Hash-Wert eines bestimmten Schlüssels mit Hilfe des
	Modulo-Operators.

	Die Funktion Insert fügt ein neues Schlüssel-Wert-Paar in die Hash-Tabelle ein. Sie
	berechnet zunächst den Hash-Wert des Schlüssels und durchläuft dann die verknüpfte Liste
	bei diesem Index, um zu prüfen, ob der Schlüssel bereits existiert. Ist dies der Fall,
	wird der Wert aktualisiert. Ist dies nicht der Fall, wird ein neuer Knoten erstellt und
	an den Anfang der verknüpften Liste gesetzt.

	Die Funktion Get ruft den einem bestimmten Schlüssel zugeordneten Wert aus der Hashtabelle
	ab. Sie berechnet zunächst den Hash-Wert des Schlüssels und durchläuft dann die verknüpfte
	Liste bei diesem Index, um den Knoten mit dem passenden Schlüssel zu finden. Wird er gefunden,
	gibt er den zugehörigen Wert zurück. Ist dies nicht der Fall, gibt sie false zurück.

	Schließlich erstelle ich in der Hauptfunktion eine neue Hashtabelle mit einer Größe von 10,
	füge drei Schlüssel-Wert-Paare ein und rufe den Wert ab, der dem Schlüssel 2 zugeordnet ist.
	Das Programm gibt zwei Werte auf der Konsole aus.

*/

type Node struct {
	Key   int
	Value string
	Next  *Node
}

type HashTable struct {
	Size  int
	Table []*Node
}

func NewHashTable(size int) *HashTable {
	table := make([]*Node, size)
	return &HashTable{size, table}
}

func (ht *HashTable) hash(key int) int {
	return key % ht.Size
}

func (ht *HashTable) Insert(key int, value string) {
	index := ht.hash(key)
	node := ht.Table[index]
	for node != nil {
		if node.Key == key {
			node.Value = value
			return
		}
		node = node.Next
	}
	newNode := &Node{Key: key, Value: value, Next: ht.Table[index]}
	ht.Table[index] = newNode
}

func (ht *HashTable) Get(key int) (string, bool) {
	index := ht.hash(key)
	node := ht.Table[index]
	for node != nil {
		if node.Key == key {
			return node.Value, true
		}
		node = node.Next
	}
	return "", false
}

func main() {

	hashTable := NewHashTable(10)
	hashTable.Insert(1, "one")
	hashTable.Insert(2, "two")
	hashTable.Insert(11, "eleven")
	val, ok := hashTable.Get(2)
	if ok {
		fmt.Println(val)
	}

}
