package main

import (
	"log"
	"math"
)

/*

	Beispiel:

	Berechnung der euklidischen Distanz zwischen zwei Punkten

*/

type Point struct {
	x float64
	y float64
}

type Line struct {
	p1 Point
	p2 Point
}

func distance(l1 Line) float64 {
	return math.Sqrt(math.Pow(l1.p2.x-l1.p1.x, 2) + math.Pow(l1.p2.y-l1.p1.y, 2))
}

func main() {

	line := Line{
		Point{float64(215), float64(237)},
		Point{float64(300), float64(237)},
	}

	log.Printf("main: euclidean distance is %.2f\n", distance(line))

}
