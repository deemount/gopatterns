package main

import "fmt"

/*
	Angenommen, ich baue einen In-Memory-Cache. Da er sich im Speicher befindet, hat er eine begrenzte Größe.
	Immer wenn er seine maximale Größe erreicht, müssen einige Einträge entfernt werden, um Platz freizugeben.
	Dies kann durch verschiedene Algorithmen geschehen.

	Einige der gängigen Algorithmen sind:

	* Least Recently Used (LRU): Der Eintrag, der am wenigsten benutzt wurde, wird entfernt.
	* First In, First Out (FIFO): Entfernen eines Eintrags, der zuerst erstellt wurde.
	* Least Frequently Used (LFU): Entfernen eines Eintrags, der am wenigsten häufig benutzt wurde.

	Das Problem besteht darin, die Cache-Klasse von diesen Algorithmen zu entkoppeln, so dass ich den
	Algorithmus zur Laufzeit ändern kann. Außerdem sollte sich die Cacheklasse nicht ändern, wenn ein
	neuer Algorithmus hinzugefügt wird.

	Hier kommt das Strategy-Muster ins Spiel. Es schlägt vor, eine Familie von Algorithmen zu erstellen,
	wobei jeder Algorithmus seine eigene Klasse hat. Jede dieser Klassen hat dieselbe Schnittstelle. Diese
	macht den Algorithmus innerhalb der Familie austauschbar. Nehmen wir an, der Name der gemeinsamen
	Schnittstelle lautet evictionAlgo.

	Meine Haupt-Cache-Klasse wird nun die Schnittstelle evictionAlgo einbetten. Anstatt alle Arten von
	Verdrängungsalgorithmen in sich selbst zu implementieren, wird meine Cache-Klasse die Ausführung an
	die Schnittstelle evictionAlgo delegieren.

	Da evictionAlgo eine Schnittstelle ist, können wir den Algorithmus zur Laufzeit entweder auf LRU, FIFO
	oder LFU ändern, ohne die Cacheklasse zu ändern.
*/

// Strategy Interface

type EvictionAlgo interface {
	evict(c *Cache)
}

// Concrete strategy

type Fifo struct{}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

// Context

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Client code

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
