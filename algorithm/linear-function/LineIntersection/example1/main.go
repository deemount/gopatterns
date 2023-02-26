package main

import (
	"fmt"
	"math"
)

/*

	Beispiel:

	Schnittpunktberechnung für Linien in einem zwei-dimensionalen Bereich von x, y: +Inf

*/

type Point struct {
	x, y float64
}

func fixError(n float64) float64 {
	return math.Round(n*100000000000) / 100000000000
}

// toFixed formats a number using fixed-point notation.
func toFixed(n float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))
	return math.Round(n*scale) / scale
}

func lineIntersect(p1, p2, p3, p4 Point) (Point, error) {

	if math.Max(p1.x, p2.x) < math.Min(p3.x, p4.x) ||
		math.Min(p1.x, p2.x) > math.Max(p3.x, p4.x) ||
		math.Max(p1.y, p2.y) < math.Min(p3.y, p4.y) ||
		math.Min(p1.y, p2.y) > math.Max(p3.y, p4.y) {
		return Point{}, fmt.Errorf("lineIntersect: point relations not given")
	}

	// calculate relative location of the intersection t and u
	nt := (p1.x*p2.y-p1.y*p2.x)*(p3.x-p4.x) - (p1.x-p2.x)*(p3.x*p4.y-p3.y*p4.x)
	nu := (p1.x*p2.y-p1.y*p2.x)*(p3.y-p4.y) - (p1.y-p2.y)*(p3.x*p4.y-p3.y*p4.x)

	// Denomination
	denominator := (p1.x-p2.x)*(p3.y-p4.y) - (p1.y-p2.y)*(p3.x-p4.x)
	if math.IsNaN(denominator) || denominator == 0 {
		return Point{}, fmt.Errorf("lineIntersect: lines are parallel")
	}

	// fxing Inf and NaN
	px := fixError(nt / denominator)
	py := fixError(nu / denominator)
	t := +toFixed(px, 2)
	u := +toFixed(py, 2)

	if t < +toFixed(math.Min(p1.x, p2.x), 2) ||
		t > +toFixed(math.Max(p1.x, p2.x), 2) ||
		t < +toFixed(math.Min(p3.x, p4.x), 2) ||
		t > +toFixed(math.Max(p3.x, p4.x), 2) ||
		u < +toFixed(math.Min(p1.y, p2.y), 2) ||
		u > +toFixed(math.Max(p1.y, p2.y), 2) ||
		u < +toFixed(math.Min(p3.y, p4.y), 2) ||
		u > +toFixed(math.Max(p3.y, p4.y), 2) {
		return Point{}, fmt.Errorf("lineIntersect: no line intersection given")
	}

	return Point{x: px, y: py}, nil

}

func main() {

	p1 := Point{178, 247}
	p2 := Point{590, 247}
	p3 := Point{530, 322}
	p4 := Point{530, 160}

	intersection, err := lineIntersect(p1, p2, p3, p4)
	if err != nil {
		fmt.Printf("main: %s", err)
	} else {
		fmt.Printf("main: intersection point:  is (%.2f, %.2f)\n", intersection.x, intersection.y)
	}

}
