package main

import (
	"fmt"
	"math"
)

/*

	Beispiel:

	In diesem Beispiel stellt die Funktion f(x) die Funktion dar,
	deren Wurzel ich mit der Bisektionsmethode finden will. Die Funktion
	bisection(a, b, tol) implementiert den Bisektionsalgorithmus, wobei die
	Endpunkte a und b des Intervalls und eine Toleranz tol berücksichtigt werden.

	Der Algorithmus wird fortgesetzt, bis der Absolutwert von f(c)
	(wobei c der Mittelpunkt des Intervalls ist) kleiner als die Toleranz tol ist.

	Die Hauptfunktion ruft bisection mit einigen Anfangswerten auf und gibt die
	gefundene Wurzel aus.

*/

func f(x float64) float64 {
	return math.Pow(x, 2) - 2
}

func bisection(a, b, tol float64) float64 {

	if f(a)*f(b) >= 0 {
		fmt.Println("Error: f(a) and f(b) must have opposite signs.")
		return 0
	}

	c := (a + b) / 2

	for math.Abs(f(c)) > tol {
		if f(c)*f(a) < 0 {
			b = c
		} else {
			a = c
		}

		c = (a + b) / 2
	}

	return c
}

func main() {
	a := 1.0
	b := 2.0
	tol := 0.0001

	root := bisection(a, b, tol)
	fmt.Printf("Root: %v\n", root)
}
