package main

import "log"

/*

	Beispiel:

	In diesem Beispiel definiere ich eine Hyperplane-Struktur mit einem Normalvektor und einer Vorspannung.
	Wir definieren auch eine Methode, um die Position eines Punktes relativ zur Hyperebene zu bewerten, indem ich
	das Punktprodukt des Punktes und des Normalvektors nehme, den Bias hinzufüge und das Ergebnis zurückgebe.
	Schließlich erstelle ich eine Hyperebene im 3D-Raum und werte einen Punkt relativ zu ihr aus. Die Ausgabe
	sollte die Position des Punktes relativ zur Hyperebene anzeigen.

*/

// Define a hyperplane struct
type Hyperplane struct {
	normal []float64 // the normal vector to the hyperplane
	bias   float64   // the bias or intercept of the hyperplane
}

// Define a method to evaluate a point's position relative to the hyperplane
func (h Hyperplane) Evaluate(point []float64) float64 {
	dotProduct := 0.0
	for i := 0; i < len(point); i++ {
		dotProduct += point[i] * h.normal[i]
	}
	return dotProduct + h.bias
}

func main() {
	// Create a hyperplane in 3D space with normal (1, 2, -3) and bias -5
	hyperplane := Hyperplane{[]float64{1, 2, -3}, -5}

	// Evaluate a point (2, 3, 4) relative to the hyperplane
	point := []float64{2, 3, 4}
	position := hyperplane.Evaluate(point)

	log.Println("The point", point, "is", position, "units away from the hyperplane.")
}
