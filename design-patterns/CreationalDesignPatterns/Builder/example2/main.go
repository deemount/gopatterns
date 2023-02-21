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

/*

 * Der IglooBuilder ist die Implementierung der Builder-Schnittstelle.
 * Die Methoden setWindowType, setDoorType und setNumFloor sind Bestandteile des Hauses.
 * Die Methode getHouse dient dazu, das gebaute Haus zurückzugeben.

 */

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

/*

* Der UsualBuilder ist die Implementierung der Builder-Schnittstelle.
* Die Methoden setWindowType, setDoorType und setNumFloor sind Bestandteile des Hauses.
* Die Methode getHouse dient dazu, das gebaute Haus zurückzugeben.

 */

type UsualBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *UsualBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *UsualBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *UsualBuilder) setNumFloor() {
	b.floor = 2
}

func (b *UsualBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func main() {

}
