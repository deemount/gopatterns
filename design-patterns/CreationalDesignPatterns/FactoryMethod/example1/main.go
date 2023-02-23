package main

import "log"

/*
 * Single Responsibility Principle (S)
 * "One class should have one, and only one, responsibility."
 */

type VerkaufSchnittstelle interface {
	berechnen() float64
}

/*
 * Open/Closed Principle (O)
 * "You should be able to extend a class's behavior without modifying it."
 */

type Verkauf struct {
	VerkaufSchnittstelle []VerkaufSchnittstelle
	gebuehr              float64
}

func (v *Verkauf) berechnen() float64 {
	for _, f := range v.VerkaufSchnittstelle {
		v.gebuehr += f.berechnen()
	}
	return v.gebuehr
}

func (v *Verkauf) hinzufuegen(f VerkaufSchnittstelle) {
	v.VerkaufSchnittstelle = append(v.VerkaufSchnittstelle, f)
}

func (v Verkauf) total() float64 {
	return v.gebuehr
}

/*
 * Interface Segregation Principle (I)
 * "Many client-specific interfaces are better than one general-purpose interface."
 */
type Einzelhandel interface {
	VerkaufSchnittstelle
}

type einzelhandel struct {
	Verkauf
}

func (v einzelhandel) berechnen() float64 {
	v.gebuehr = 1.1
	return v.gebuehr
}

type Aussenhandel interface {
	VerkaufSchnittstelle
}

type aussenhandel struct {
	Verkauf
}

func (v aussenhandel) berechnen() float64 {
	v.gebuehr = 2.0
	return v.gebuehr
}

/*
 * Factory Method
 * A factory is an object that is used to create other objects.
 * In a factory method pattern, a helper method (or function) is defined,
 * to enable object creation without knowing the implementation class details.
 */

func NeuVerkauf(typ string) VerkaufSchnittstelle {
	switch typ {
	case "Einzelhandel":
		return einzelhandel{}
	case "Aussenhandel":
		return aussenhandel{}
	default:
		return nil
	}
}

// Execute

func main() {

	verkauf := new(Verkauf)

	verkauf.hinzufuegen(NeuVerkauf("Einzelhandel"))
	verkauf.berechnen()

	verkauf.hinzufuegen(NeuVerkauf("Aussenhandel"))
	verkauf.berechnen()

	log.Printf("%#v", verkauf)
	log.Printf("main: gebühren (total) %#v", verkauf.total())

}
