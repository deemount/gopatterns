package main

import (
	"fmt"
	"strings"
)

/*

	Beispiel:

	Diese Implementierung verwendet eine Scheibe von Runen, um den Puffer zu speichern
	und behält die Lücke zwischen gapStart und gapEnd im Auge. Die Insert-Methode
	vergrößert die Lücke bei Bedarf, kopiert den Text in die Lücke und passt den
	gapStart-Zeiger an. Die Delete-Methode kopiert den Text nach der Lücke in die Lücke
	und passt den GapEnd-Zeiger an. Die expandGap-Methode erstellt einen neuen Puffer mit
	genügend Platz für die Lücke, kopiert den Text vor und nach der Lücke in den neuen Puffer
	und aktualisiert den gapEnd-Zeiger. Schließlich gibt die String-Methode den Puffer zurück,
	in dem die Lücke entfernt wurde.

*/

type GapBuffer struct {
	buffer   []rune
	gapStart int
	gapEnd   int
}

func (g *GapBuffer) Insert(text string) {
	gapLength := g.gapEnd - g.gapStart
	textLength := len(text)
	if gapLength < textLength {
		g.expandGap(textLength - gapLength)
	}
	copy(g.buffer[g.gapStart+textLength:], g.buffer[g.gapEnd:])
	copy(g.buffer[g.gapStart:], []rune(text))
	g.gapStart += textLength
}

func (g *GapBuffer) Delete(length int) {
	copy(g.buffer[g.gapStart:], g.buffer[g.gapEnd:g.gapEnd+length])
	g.gapEnd += length
}

func (g *GapBuffer) expandGap(length int) {
	newBuffer := make([]rune, len(g.buffer)+length)
	copy(newBuffer[:g.gapStart], g.buffer[:g.gapStart])
	copy(newBuffer[g.gapStart+length:], g.buffer[g.gapStart:])
	g.buffer = newBuffer
	g.gapEnd += length
}

func (g *GapBuffer) String() string {
	return strings.ReplaceAll(string(g.buffer), string(g.buffer[g.gapStart:g.gapEnd]), "")
}

func main() {
	gb := GapBuffer{make([]rune, 10), 0, 10}
	gb.Insert("Hello, ")
	gb.Insert("world!")
	fmt.Println(gb) // Output: Hello, world!
	gb.Delete(7)
	fmt.Println(gb) // Output: Hello!
}
