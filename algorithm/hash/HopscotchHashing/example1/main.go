package main

import "fmt"

/*

	Beispiel:

	Diese Implementierung definiert eine HopscotchTable-Struktur mit einem Array
	von Schlüsseln und einem Array von Werten (in diesem Fall Ganzzahlen) und
	implementiert Methoden zum Hashing, zum Auffinden eines leeren Slots,
	zum Auffinden eines Schlüssels und zum Einfügen eines Schlüssel-Wert-Paares.

	Die insert-Methode implementiert den hopscotch-Algorithmus, indem sie durch einen
	maximalen Abstand maxDist iteriert und Elemente im Array austauscht, bis ein
	leerer Slot gefunden oder der maximale Abstand überschritten wird.

	In der Hauptfunktion erstelle ich eine HopscotchTable-Instanz h, füge einige
	Schlüssel-Wert-Paare ein und gebe die resultierende Tabelle und einige
	Beispielergebnisse aus.

*/

const (
	tableSize = 16 // Size of hash table
	maxDist   = 4  // Maximum distance for hopscotch algorithm
)

type HopscotchTable struct {
	keys  [tableSize]int
	table [tableSize]int
}

func (h *HopscotchTable) hash(key int) int {
	return key % tableSize
}

func (h *HopscotchTable) findEmpty(start int) int {
	for i := start; i < tableSize; i++ {
		if h.keys[i] == 0 {
			return i
		}
	}
	for i := 0; i < start; i++ {
		if h.keys[i] == 0 {
			return i
		}
	}
	return -1
}

func (h *HopscotchTable) find(key int) (int, bool) {
	hash := h.hash(key)
	for i := 0; i <= maxDist; i++ {
		index := (hash + i) % tableSize
		if h.keys[index] == key {
			return index, true
		}
	}
	return -1, false
}

func (h *HopscotchTable) insert(key, value int) bool {
	_, found := h.find(key)
	if found {
		return false
	}
	emptyIndex := h.findEmpty(0)
	if emptyIndex == -1 {
		return false
	}
	for i := 0; i < maxDist; i++ {
		currIndex := (emptyIndex + i) % tableSize
		if i > 0 && currIndex == emptyIndex {
			return false // Cycle detected, cannot insert key
		}
		if h.keys[currIndex] == 0 {
			emptyIndex = currIndex
		} else {
			hopIndex := (emptyIndex + maxDist) % tableSize
			h.keys[emptyIndex] = key
			h.table[emptyIndex] = value
			h.keys[hopIndex], h.table[hopIndex] = h.keys[currIndex], h.table[currIndex]
			emptyIndex = hopIndex
			i = -1 // Start over from beginning of maxDist
		}
	}
	h.keys[emptyIndex] = key
	h.table[emptyIndex] = value
	return true
}

func main() {
	h := HopscotchTable{}
	h.insert(1, 10)
	h.insert(2, 20)
	h.insert(3, 30)
	h.insert(17, 40)
	h.insert(18, 50)
	fmt.Println(h.table)
	fmt.Println(h.find(1))
	fmt.Println(h.find(18))
	fmt.Println(h.find(100))
}
