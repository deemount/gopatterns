package main

/*
	* Builder ist die Schnittstelle, die die Methoden zur Erstellung der Produktteile festlegt.
	* Das Haus stellt das komplexe Objekt dar, das wir bauen wollen.
	* Der Direktor ist für die Verwendung der Builder-Schnittstelle zur Herstellung des Produkts verantwortlich.
	* Mit setBuilder kann der Builder, der vom Director verwendet wird, geändert werden.
	* Der BuildHouse verwendet den gespeicherten Builder, um die erforderlichen Schritte zu reproduzieren und
	  gibt das gebaute Haus zurück.
*/

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Director struct {
	builder Builder
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()

	return d.builder.getHouse()
}
