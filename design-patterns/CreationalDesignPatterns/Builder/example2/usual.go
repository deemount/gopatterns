package main

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
