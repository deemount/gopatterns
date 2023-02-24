package main

import (
	"container/list"
	"fmt"
)

/*

	Beispiel:

	In dieser Implementierung ist MRUCache eine Struktur,
		* die die Kapazität des Caches enthält,
		* eine Karte, die die Cache-Einträge speichert (verschlüsselt durch den Schlüssel) und
		* eine  doppelt verknüpfte Liste, die die Reihenfolge der Cache-Einträge darstellt
		  (mit dem zuletzt verwendeten Eintrag an der Spitze der Liste).

	Get nimmt einen Schlüssel und gibt den entsprechenden Wert zurück, wenn der Schlüssel
	im Cache vorhanden ist. Wenn der Schlüssel gefunden wird, wird das entsprechende
	Listenelement an den Anfang der Liste verschoben (da es nun der zuletzt verwendete Eintrag ist).
	Wenn der Schlüssel nicht gefunden wird, wird -1 zurückgegeben.

	Put nimmt einen Schlüssel und einen Wert und fügt sie dem Cache hinzu. Wenn der Schlüssel bereits
	im Cache vorhanden ist, wird der entsprechende Wert aktualisiert und das Listenelement an den Anfang
	der Liste verschoben. Wenn der Schlüssel nicht im Cache vorhanden ist und der Cache bereits voll ist,
	wird der zuletzt verwendete Eintrag aus der Liste und dem Cache entfernt.
	Anschließend wird ein neues Listenelement erstellt und an den Anfang der Liste gesetzt und der
	entsprechende Schlüssel und das Element werden dem Cache hinzugefügt.

*/

type MRUCache struct {
	capacity int
	cache    map[int]*list.Element
	queue    *list.List
}

type Entry struct {
	key   int
	value int
}

func NewMRUCache(capacity int) *MRUCache {
	return &MRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (cache *MRUCache) Get(key int) int {
	if elem, ok := cache.cache[key]; ok {
		cache.queue.MoveToFront(elem)
		return elem.Value.(*Entry).value
	}
	return -1
}

func (cache *MRUCache) Put(key int, value int) {
	if elem, ok := cache.cache[key]; ok {
		elem.Value.(*Entry).value = value
		cache.queue.MoveToFront(elem)
	} else {
		if cache.queue.Len() == cache.capacity {
			lastElem := cache.queue.Back()
			delete(cache.cache, lastElem.Value.(*Entry).key)
			cache.queue.Remove(lastElem)
		}
		newElem := cache.queue.PushFront(&Entry{key: key, value: value})
		cache.cache[key] = newElem
	}
}

func main() {
	cache := NewMRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // Output: -1
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
