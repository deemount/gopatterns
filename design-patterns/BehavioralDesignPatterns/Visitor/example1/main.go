package main

import "fmt"

/*
	Beispiel:

	Nehmen wir an, ich bin der Betreuer einer Bibliothek, die verschiedene Form-Strukturen hat, wie z.B.:

	* Quadrat
	* Kreis
	* Dreieck

	Jede der oben genannten Formstrukturen implementiert die gemeinsame Formschnittstelle.
	Als die Leute in meiner Firma anfingen, meine großartige Bibliothek zu benutzen,
	wurde ich mit Funktionsanfragen überflutet. Ein Team bat mich, das getArea-Verhalten
	zu den Form-Strukturen hinzuzufügen.

	Es gibt viele Möglichkeiten, dieses Problem zu lösen:

	* Die erste Option, ist die getArea-Methode direkt in die Form-Schnittstelle einzufügen und sie dann
	  in jeder Form-Struktur zu implementieren.

	Dies scheint eine gute Lösung zu sein, aber sie hat ihren Preis. Als Betreuer der Bibliothek
	möchte ich nicht riskieren, den wertvollen Code jedes Mal zu zerstören, wenn jemand nach
	einem anderen Verhalten fragt.

	Dennoch möchte ich, dass andere Teams meine Bibliothek irgendwie erweitern.

	* Die zweite Möglichkeit ist, dass das Team, das die Funktion anfordert, das Verhalten selbst
	  implementieren kann.

    Dies ist jedoch nicht immer möglich, da dieses Verhalten von meinem
	privaten Code abhängen kann.

	* Die dritte Möglichkeit besteht darin, das obige Problem mit Hilfe des
	  Visitor-Entwurfsmuster zu lösen.

	Ich beginne also mit der Definition einer sog. Visitor-Schnittstelle:

*/

// Visitor

type Visitor interface {
	besucherQuadrat(*Quadrat)
	besucherKreis(*Kreis)
	besucherRechteck(*Rechteck)
}

// Concrete Visitor

type Flaechenberechnung struct {
	area int
}

func (a *Flaechenberechnung) besucherQuadrat(s *Quadrat) {
	// Calculate area for Quadrat.
	// Then assign in to the area instance variable.
	fmt.Println("berechnung der fläche für quadrat")
}

func (a *Flaechenberechnung) besucherKreis(s *Kreis) {
	fmt.Println("berechnung der fläche für kreis")
}
func (a *Flaechenberechnung) besucherRechteck(s *Rechteck) {
	fmt.Println("berechnung der fläche für rechteck")
}

type MittelpunktKoordinaten struct {
	x int
	y int
}

func (a *MittelpunktKoordinaten) besucherQuadrat(s *Quadrat) {
	fmt.Println("berechnung der koordinaten des mittelpunkts für quadrat")
}

func (a *MittelpunktKoordinaten) besucherKreis(c *Kreis) {
	fmt.Println("berechnung der koordinaten des mittelpunkts für kreis")
}
func (a *MittelpunktKoordinaten) besucherRechteck(t *Rechteck) {
	fmt.Println("berechnung der koordinaten des mittelpunkts für rechteck")
}

/*
	Die Funktionen besucherQuadrat(Quadrat), besucherKreis(Kreis), besucherTriangle(Dreieck)
	ermöglichen es mir, Quadraten, Kreisen bzw. Dreiecken Funktionalität hinzuzufügen.

	Der zweite wichtige Teil ist das Hinzufügen der akzeptiere-Methode zur shape-Schnittstelle.

	func akzeptiere(v visitor)

	Alle Form-Strukturen müssen diese Methode definieren, zum Beispiel:

	func (obj *Quadrat) akzeptiere(v visitor) {
		v.besucherQuadrat(obj)
	}

	Wenn ich das Visitor-Muster verwende, muss ich dennoch unsere Form-Strukturen ändern.
	Diese Änderung erfolgt nur einmal!

	Wenn ich andere Verhaltensweisen wie getNumSides, getMittelpunktKoordinaten hinzufüge,
	werde ich die gleiche Funktion akzeptiere(v visitor) ohne weitere Änderungen an den
	Form-Strukturen verwenden.

	Am Ende müssen die Form-Strukturen nur einmal geändert werden. Alle zukünftigen
	Anfragen für verschiedene Verhaltensweisen können mit der gleichen akzeptiere-Funktion
	behandelt werden.

	Wenn nun das Team das getArea-Verhalten anfordert, kann ich einfach
	die konkrete Implementierung der Besucherschnittstelle definieren und die Logik der
	Flächenberechnung Logik in diese konkrete Implementierung schreiben.

*/

// Element

type Form interface {
	typ() string
	akzeptiere(Visitor)
}

// Concrete Element

type Quadrat struct {
	side int
}

func (s *Quadrat) akzeptiere(v Visitor) {
	v.besucherQuadrat(s)
}

func (s *Quadrat) typ() string {
	return "Quadrat"
}

type Kreis struct {
	radius int
}

func (c *Kreis) akzeptiere(v Visitor) {
	v.besucherKreis(c)
}

func (c *Kreis) typ() string {
	return "Kreis"
}

type Rechteck struct {
	l int
	b int
}

func (t *Rechteck) akzeptiere(v Visitor) {
	v.besucherRechteck(t)
}

func (t *Rechteck) typ() string {
	return "Rechteck"
}

// Client code

func main() {

	Quadrat := &Quadrat{side: 2}
	Kreis := &Kreis{radius: 3}
	Rechteck := &Rechteck{l: 2, b: 3}

	flaeche := &Flaechenberechnung{}

	Quadrat.akzeptiere(flaeche)
	Kreis.akzeptiere(flaeche)
	Rechteck.akzeptiere(flaeche)

	fmt.Println()

	koordinaten := &MittelpunktKoordinaten{}
	Quadrat.akzeptiere(koordinaten)
	Kreis.akzeptiere(koordinaten)
	Rechteck.akzeptiere(koordinaten)
}
