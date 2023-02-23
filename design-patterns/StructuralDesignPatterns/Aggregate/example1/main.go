package main

import "fmt"

/*

	Beispiel:

	Ich definiere einen benutzerdefinierten Typ Item, der einzelne Objekte darstellt.
	Dazu einen benutzerdefinierten Typ Catalog, der eine Sammlung von Item-Objekten darstellt.
	Ich definieren zudem eine Schnittstelle Iterator, um den Zugriff auf die einzelnen Objekte
	in der Sammlung zu ermöglichen und eine Implementierung der Iterator-Schnittstelle für den
	Typ Catalog namens CatalogIterator.

	Der Typ Catalog verfügt über eine Methode GetIterator(), die ein Iterator-Objekt zurückgibt,
	mit dem Clients über die Sammlung von Item-Objekten iterieren können. In der Funktion main()
	erstelle ich einige Item-Objekte und füge diese zu einem Catalog-Objekt hinzu.

	Anschließend verwende ich die Iterator-Schnittstelle, um über das Catalog-Objekt zu iterieren
	und Informationen über jedes Item-Objekt auszudrucken.

	Diese Implementierung des Aggregate-Patterns ermöglicht es den Clients, die Sammlung von Item-Objekten
	und die einzelnen Item-Objekte einheitlich zu behandeln, was die Arbeit mit Sammlungen von Objekten
	auf konsistente und vorhersehbare Weise erleichtert.

*/

// Define a custom type for the individual objects
type Item struct {
	Name  string
	Price float64
}

// Define a custom type for the collection of objects
type Catalog struct {
	Items []*Item
}

// Define an interface to provide access to the individual objects in the collection
type Iterator interface {
	Next() *Item
	HasNext() bool
}

// Implement the Iterator interface for the Catalog type
type CatalogIterator struct {
	catalog *Catalog
	index   int
}

func (ci *CatalogIterator) Next() *Item {
	item := ci.catalog.Items[ci.index]
	ci.index++
	return item
}

func (ci *CatalogIterator) HasNext() bool {
	return ci.index < len(ci.catalog.Items)
}

func (c *Catalog) GetIterator() Iterator {
	return &CatalogIterator{catalog: c, index: 0}
}

func main() {
	// Create some individual objects
	item1 := &Item{Name: "Item 1", Price: 10.00}
	item2 := &Item{Name: "Item 2", Price: 20.00}
	item3 := &Item{Name: "Item 3", Price: 30.00}

	// Create the collection object
	catalog := &Catalog{Items: []*Item{item1, item2, item3}}

	// Use the Iterator interface to iterate over the collection of objects
	iterator := catalog.GetIterator()
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Printf("%s: $%.2f\n", item.Name, item.Price)
	}
}
