package main

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
