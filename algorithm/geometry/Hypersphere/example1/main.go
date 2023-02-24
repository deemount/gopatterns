package main

import (
	"log"
	"math"
)

/*

	Beispiel:

	In diesem Beispiel definiere ich eine Point-Struktur, um einen Punkt
	im n-dimensionalen Raum darzustellen und eine Hypersphere-Struktur,
	um eine Hypersphäre darzustellen, deren Mittelpunkt ein bestimmter Punkt
	mit einem bestimmten Radius ist. Anschließend definiere ich Methoden für
	die Point- und Hypersphere-Strukturen, um den Abstand zwischen zwei Punkten
	zu berechnen und um festzustellen, ob ein bestimmter Punkt in der Hypersphäre
	enthalten ist.

	In der Hauptfunktion erstelle ich eine 3-dimensionale Hypersphäre mit dem
	Zentrum im Ursprung und dem Radius 1. Anschließend wird mit der Contains-Methode
	geprüft, ob zwei Punkte innerhalb der Hypersphäre enthalten sind. Der erste Punkt
	befindet sich innerhalb der Hypersphäre, der zweite außerhalb, so dass die Ausgabe
	true bzw. false ist.

*/

type Point struct {
	coords []float64
}

func NewPoint(coords []float64) *Point {
	return &Point{coords}
}

func (p *Point) DistanceTo(other *Point) float64 {
	sum := 0.0
	for i, coord := range p.coords {
		sum += math.Pow(other.coords[i]-coord, 2)
	}
	return math.Sqrt(sum)
}

type Hypersphere struct {
	center *Point
	radius float64
}

func NewHypersphere(center *Point, radius float64) *Hypersphere {
	return &Hypersphere{center, radius}
}

func (s *Hypersphere) Contains(p *Point) bool {
	return s.center.DistanceTo(p) <= s.radius
}

func main() {

	center := NewPoint([]float64{0, 0, 0}) // center of hypersphere
	radius := 1.0                          // radius of hypersphere
	sphere := NewHypersphere(center, radius)

	pointInside := NewPoint([]float64{0.5, 0.5, 0.5})
	pointOutside := NewPoint([]float64{1, 1, 1})

	log.Println(sphere.Contains(pointInside))  // true
	log.Println(sphere.Contains(pointOutside)) // false

}
