package main

import (
	"container/list"
	"fmt"
	"log"
)

/*

	Beispiel:

	Ich muss die erforderlichen Typen modellieren, die das Rückgrat meines Cache bilden werden.

	Die erste Struktur ist das CacheItem. Dies wird das eigentliche Element sein, das im Cache
	gespeichert wird. Es enthält den Schlüssel, mit dem ich in der Hash-Tabelle den Wert nachschlagen kann,
	der das eigentliche gecachte Element darstellt, und einen frequencyParent-Zeiger auf den Zeiger
	in der Frequenzliste.

	Die nächste Struktur ist FrequencyItem, die jedes Element der Häufigkeitsliste darstellt. Es enthält
	eine Menge von Einträgen, die eine Menge von CacheItem-Zeigern sein wird. Ich werde eine Map verwenden,
	um sie zu speichern, damit ich sie als eine Menge behandeln kann, die nur eindeutige Elemente enthält.

	Die letzte Struktur, die ich brauchen, um einen reibungslos funktionierenden Cache zu haben,
	ist der Cache selbst.

	Der Cache enthält
		* die Hash-Map, genannt bykey,
		* die Frequenzliste, genannt freqs,
		* die maximale Kapazität des Cache, genannt capacity und
		* die Größe des Cache, die die Anzahl der zu einem bestimmten Zeitpunkt
		  zwischengespeicherten Elemente darstellt.

	Der Konstruktor New erzeugt eine neue Cache-Struktur und setzt alle Standardwerte. Für die Frequenzliste
	list.New() verwenden wir das Container/List-Paket von Go, das eine ordentliche Implementierung von
	verknüpften Listen enthält.

	Die Set-Funktion erhält als Argumente den Cache-Schlüssel und den zu cachenden Wert/Eintrag.
	Dann prüft sie, ob das Element bereits im Cache gespeichert ist oder nicht. Wenn es zwischengespeichert
	ist, wird nur der Wert des Elements aktualisiert. Andernfalls wird ein neues CacheItem erstellt, das den
	aktuellen Wert kapselt, den Schlüssel festlegt, das Element zur Hash-Tabelle bykey hinzufügt und die
	Größe des Cache erhöht.

	* Der Cache muss wissen, wie er die Zugriffshäufigkeit für aCacheItem erhöhen kann
	* Der Cache muss wissen, wie er ein Element auf der Grundlage der Zugriffshäufigkeit
	  entfernen kann, wenn die Größe die Kapazität erreicht

	Die dritte Funktion, die Cache erhält, ist die Get-methode, die Zugriff auf das Element
	anhand des Schlüssels aus der Hashtabelle hat und diesen Wert ausgibt. Ich prüfe also,
	ob die Hashtabelle bykey einen Wert mit dem Argument key enthält und falls ja, wird dieser
	zurückgegeben. Andernfalls gebe ich nil zurück. Wenn ich in der Get-Funktion ein Element in der
	Bykey-Hashtabelle finde, muss ich seine Zugriffshäufigkeit erhöhen, bevor ich seinen Wert
	zurückgeben kann. Der Cache wird diese Häufigkeit dieses Elements erhöhen.

	Sobald ich den Knoten in der Liste kenne, aus dem ich ihn entfernen will, kann ich das
	Element einfach aus der Liste der Einträge entfernen und auch das FrequencyItem vollständig
	aus der Frequenzliste entfernen, wenn die Einträge leer werden.

	Aktualisieren der Zugriffshäufigkeit

	Ich muss bei jedem Zugriff auf den Cache die Zugriffshäufigkeit des aufgerufenen Elements
	aktualisieren.

	Ich schaue mir nun die Schritte an, die meine Increment-Funktion durchführen muss. Zunächst
	müssen wir für das ablaufende Element entscheiden, ob dieses Element bereits Teil der Hashtabelle
	und der Häufigkeitsliste ist oder nicht. Wenn ja, muss ich seinen neuen Häufigkeitswert und
	seine nächste Häufigkeitsposition (Knoten) in der Häufigkeitsliste herausfinden.

	Zweitens muss ich herausfinden, ob es für die neue Frequenz bereits einen Knoten in der Frequenzliste
	gibt oder nicht. Ist dies der Fall, muss ich den Eintrag in die Liste aufnehmen und ihm die neue
	Zugangsfrequenz zuweisen (die der aktuellen Zugangsfrequenz + 1 entspricht). Wenn es keinen gibt,
	muss ich einen neuen Frequenzknoten in der Frequenzliste erstellen (und alle seine vernünftigen
	Standardwerte einstellen) und dann das Element zu seiner Liste von Einträgen hinzufügen.

	Drittens muss die Funktion, sobald ich einen FrequencyParent gefunden habe, den neuen Parent auf
	das Element, das inkrementiert werden soll, setzen und es der Liste der Einträge des Parents hinzufügen.

	Als letzten Schritt entfernt die Inkrement-Funktion das Element aus den Einträgen des alten
	Frequenzknotens (frequencyParent).

	Das letzte Teil des Puzzles ist das Auslagern, d. h. das Entfernen der am wenigsten genutzten
	Elemente, sobald der Cache seine maximale Kapazität erreicht hat.

	Ich muss wissen, wie viele Elemente wir aus dem Cache entfernen wollen. Normalerweise hat unser
	Cache eine untere und eine obere Grenze. Wenn die obere Grenze erreicht ist, entfernen wir die
	Elemente bis zur unteren Grenze. In unserem Fall werden wir eine beliebige Anzahl von Objekten aus
	dem Cache entfernen, die der Funktion Evict als Argument übergeben wird.

	Die Funktion beginnt damit, die Häufigkeitsliste vom Anfang zum Ende hin "durchzugehen". Da die
	Häufigkeitsliste in aufsteigender Reihenfolge ist, beginnt sie mit dem Entfernen der Einträge ab
	dem ersten Häufigkeitsknoten, bis sie so viele Einträge entfernt hat, wie sie als beliebige Zahl
	übergeben hat.

	Wenn ein Frequenzknoten keine Einträge mehr enthält, muss die Funktion Evict diesen Frequenzknoten
	ebenfalls aus der Frequenzliste entfernen. Dazu ruft sie die Funktion Remove auf. Auf diese Weise
	hinterlässt die Verdrängung keinen Müll.

*/

type CacheItem struct {
	key             string        // Key of entry
	value           interface{}   // Value of item
	frequencyParent *list.Element // Pointer to parent in cacheList
}

type FrequencyItem struct {
	entries map[*CacheItem]byte // Set of entries
	freq    int                 // Access frequency
}

type Cache struct {
	bykey    map[string]*CacheItem // Hashmap containing *CacheItems for O(1) access
	freqs    *list.List            // Linked list of frequencies
	capacity int                   // Max number of items
	size     int                   // Current size of cache
}

func New() *Cache {
	cache := new(Cache)
	cache.bykey = make(map[string]*CacheItem)
	cache.freqs = list.New()
	cache.size = 0
	cache.capacity = 100

	return cache
}

func (cache *Cache) Set(key string, value interface{}) {
	if item, ok := cache.bykey[key]; ok {
		item.value = value
		cache.increment(item)
	} else {
		item := new(CacheItem)
		item.key = key
		item.value = value
		cache.bykey[key] = item
		cache.size++
		// For simplicity, remove only 10 items every time it reaches max capacity
		if cache.atCapacity() {
			cache.Evict(10)
		}
		// Eviction, if needed
		cache.increment(item)
	}
}

func (cache *Cache) Get(key string) interface{} {
	if e, ok := cache.bykey[key]; ok {
		cache.increment(e)
		return e.value
	}
	return nil
}

func (cache *Cache) increment(item *CacheItem) {

	currentFrequency := item.frequencyParent

	var nextFrequencyAmount int
	var nextFrequency *list.Element

	if currentFrequency == nil {
		log.Println("increment: current frequency is nil")
		nextFrequencyAmount = 1
		nextFrequency = cache.freqs.Front()
	} else {
		log.Println("increment: current frequency next")
		nextFrequencyAmount = currentFrequency.Value.(*FrequencyItem).freq + 1
		nextFrequency = currentFrequency.Next()
	}

	if nextFrequency == nil || nextFrequency.Value.(*FrequencyItem).freq != nextFrequencyAmount {
		log.Println("increment: next frequency is nil or has not the same amount")

		newFrequencyItem := new(FrequencyItem)
		newFrequencyItem.freq = nextFrequencyAmount
		newFrequencyItem.entries = make(map[*CacheItem]byte)

		if currentFrequency == nil {
			nextFrequency = cache.freqs.PushFront(newFrequencyItem)
		} else {
			nextFrequency = cache.freqs.InsertAfter(newFrequencyItem, currentFrequency)
		}

	}

	item.frequencyParent = nextFrequency
	nextFrequency.Value.(*FrequencyItem).entries[item] = 1
	if currentFrequency != nil {
		cache.Remove(currentFrequency, item)
	}
}

func (cache *Cache) Remove(listItem *list.Element, item *CacheItem) {
	frequencyItem := listItem.Value.(*FrequencyItem)
	delete(frequencyItem.entries, item)
	if len(frequencyItem.entries) == 0 {
		cache.freqs.Remove(listItem)
	}
}

func (cache *Cache) Evict(count int) {
	for i := 0; i < count; {
		if item := cache.freqs.Front(); item != nil {
			for entry, _ := range item.Value.(*FrequencyItem).entries {
				if i < count {
					delete(cache.bykey, entry.key)
					cache.Remove(item, entry)
					log.Printf("evict: removed item, entry %v, %v ", item, entry)
					cache.size--
					i++
				}
			}
		}
	}
}

func (cache *Cache) atCapacity() bool {
	log.Printf("at capacity: %v", cache.size >= cache.capacity)
	return cache.size >= cache.capacity
}

func main() {
	c := New()
	for i := 0; i < 100; i++ {
		c.Set(fmt.Sprintf("value%d", i), int(i))
	}
	log.Printf("main: get %v", c.Get("value3"))

	c.Evict(10)
	log.Printf("main: bykey %v", c.bykey)
}
