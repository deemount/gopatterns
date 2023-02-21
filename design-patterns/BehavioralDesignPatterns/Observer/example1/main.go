package main

import "fmt"

/*
	Auf der E-Commerce-Website sind Artikel von Zeit zu Zeit nicht mehr auf Lager.
	Es kann Kunden geben, die an einem bestimmten Artikel interessiert sind,
	der nicht mehr auf Lager ist.

	Es gibt drei Lösungen für dieses Problem:

	* Der Kunde prüft in regelmäßigen Abständen, ob der Artikel noch verfügbar ist
	* Der elektronische Handel bombardiert den Kunden mit allen neuen Artikeln, die auf Lager sind
	* Der Kunde abonniert nur den Artikel, an dem er interessiert ist und wird benachrichtigt,
	  wenn der Artikel verfügbar ist.
	* Außerdem können mehrere Kunden ein und dasselbe Produkt abonnieren.

	Option 3 ist am praktikabelsten und genau darum geht es beim Observer-Muster.

	Die Hauptkomponenten des Observer-Musters sind:

	* Subjekt: die Instanz, die ein Ereignis veröffentlicht, wenn etwas passiert.
	* Beobachter: Abonniert Subjekt-Ereignisse und wird benachrichtigt, wenn sie eintreten.
*/

// Subject

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

// Concrete Subject

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// Observer

type Observer interface {
	update(string)
	getID() string
}

// Concrete Observer

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

// Client code

func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
