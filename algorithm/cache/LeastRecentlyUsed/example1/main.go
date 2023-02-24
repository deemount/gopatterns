package main

import (
	"container/list"
	"log"
)

/*

	Beispiel:

	In dieser Implementierung enthält die LRUCache-Struktur eine Zuordnung von Schlüsseln zu
	Listenelementen und eine doppelt verkettete Liste, um die Reihenfolge der Elemente zu verfolgen.

	Wenn auf ein Element zugegriffen oder es eingefügt wird, wird es an den Anfang der Liste verschoben.
	Wenn die Kapazität erreicht ist, wird das am wenigsten genutzte Element vom Ende der Liste verdrängt.

	Die Get-Methode ruft den einem Schlüssel zugeordneten Wert ab, und wenn der Schlüssel gefunden wird,
	verschiebt sie das entsprechende Listenelement an den Anfang der Liste.

	Die Put-Methode fügt ein Schlüssel-Wert-Paar in den Cache ein oder aktualisiert es. Wenn der Schlüssel
	bereits existiert, wird der Wert aktualisiert und das entsprechende Listenelement an den Anfang der Liste
	verschoben. Wenn der Schlüssel nicht existiert und die Kapazität erreicht ist, wird das am wenigsten
	genutzte Element vom Ende der Liste entfernt und der entsprechende Schlüssel aus der Map gelöscht.
	Schließlich fügt es das neue Element am Anfang der Liste ein und fügt seinen Schlüssel zur Karte hinzu.

*/

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	lruList  *list.List
}

type cacheItem struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		lruList:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.lruList.MoveToFront(elem)
		return elem.Value.(*cacheItem).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.lruList.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
	} else {
		if this.lruList.Len() == this.capacity {
			lruElem := this.lruList.Back()
			lruItem := lruElem.Value.(*cacheItem)
			delete(this.cache, lruItem.key)
			this.lruList.Remove(lruElem)
		}
		newItem := &cacheItem{key: key, value: value}
		elem := this.lruList.PushFront(newItem)
		this.cache[key] = elem
	}
}

func main() {

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	log.Println(cache.Get(1)) // returns 1

	cache.Put(3, 3)

	log.Println(cache.Get(2)) // returns -1

	cache.Put(4, 4)

	log.Println(cache.Get(1)) // returns -1
	log.Println(cache.Get(3)) // returns 3
	log.Println(cache.Get(4)) // returns 4

}
