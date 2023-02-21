package main

import "fmt"

/*
	In einer Partie Counter-Strike haben Terrorist und Counter-Terrorist
	jeweils eine andere Art von Kleidung. Der Einfachheit halber nehme ich an,
	dass sowohl Terrorist als auch Counter-Terrorist jeweils einen Kleidungs-Typ haben.
	Das Kleidungsobjekt ist, wie unten dargestellt, in das Spielerobjekt eingebettet.

	Ich nehme an, dass es 5 Terroristen und 5 Gegenterroristen gibt, also insgesamt 10 Spieler.
	Nun gibt es zwei Optionen für die Kleidung.

	Jedes der 10 Spielerobjekte erstellt ein anderes Kleidungsobjekt und bettet es ein.
	Es werden insgesamt 10 Kleidungsobjekte erstellt.

	Ich erstelle zwei Kleidungsobjekte:

	* Einzelnes Terroristen-Kleidungsobjekt: Dieses wird auf 5 Terroristen aufgeteilt.
	* Ein einzelnes Kleidungsobjekt für den Terroristengegner: Dieses wird auf 5 Gegenterroristen aufgeteilt.

	Wie ich sehen kann, werden in Ansatz 1 insgesamt 10 Kleidungsobjekte erstellt,
	während in Ansatz 2 nur zwei Kleidungsobjekte erstellt werden. Der zweite Ansatz ist der,
	den ich im Flyweight Entwurfsmuster verwende. Die beiden Kleidungsobjekte, die ich erstellt habe,
	werden als Flyweight-Objekte bezeichnet.

	Das Flyweight-Muster nimmt die gemeinsamen Teile heraus und erstellt sog. Flyweight-Objekte.
	Diese Flyweight-Objekte (Kleid) können dann von mehreren Objekten (Spieler) gemeinsam genutzt werden.
	Dadurch wird die Anzahl der Kleidungsobjekte drastisch reduziert, und das Gute daran ist, dass selbst,
	wenn ich mehr Spieler erstelle, nur zwei Kleidungsobjekte ausreichen werden.

	Im Flyweight-Entwurfsmuster speichere ich die Flyweight-Objekte im Kartenfeld. Wann immer die
	anderen Objekte, die sich die Flyweight-Objekte teilen, werden die Flyweight-Objekte
	aus der Map geholt.

	Lass mich sehen, welche Teile dieser Anordnung intrinsische und extrinsische Zustände sein werden:

	* Intrinsischer Zustand: Kleid im intrinsischen Zustand, da er von mehreren Terroristen- und Gegenterroristen-Objekten gemeinsam genutzt werden kann.
	* Extrinsischer Zustand: Der Spielerstandort und die Spielerwaffe sind ein extrinsischer Zustand, da sie für jedes Objekt unterschiedlich sind.

	Unten sehe ich die Struktur für einen Spieler. Ich kann sehen, dass das Kleidungsobjekt in die Player-Struktur eingebettet ist
*/

type player struct {
	dress      Dress
	playerType string // Can be T or CT
	lat        int
	long       int
}

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

// Flyweight factory

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

// Flyweight interface

type Dress interface {
	getColor() string
}

// Concrete flyweight object

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

// Concrete flyweight object

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

// Context

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// Client code

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
