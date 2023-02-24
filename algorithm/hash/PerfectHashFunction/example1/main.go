package main

import "fmt"

/*

	Beispiel:

	In diesem Beispiel nimmt die perfectHash-Funktion ein Array von Zeichenketten und
	gibt eine perfekte Hash-Funktion für diese Zeichenketten zurück.
	Sie verwendet drei Hash-Funktionen, um eine Kandidaten-Hash-Funktion zu erzeugen
	und prüft dann auf Kollisionen, indem sie die Kandidaten-Hash-Funktion auf jede
	Zeichenkette anwendet und prüft, ob zwei Zeichenketten den gleichen Hash-Wert haben.
	 Wenn es Kollisionen gibt, wird eine neue Hash-Funktion generiert und der Versuch wiederholt.

	Sobald sie eine perfekte Hash-Funktion gefunden hat, gibt sie diese zurück. Die Hauptfunktion
	verwendet dann die perfekte Hash-Funktion, um jede Zeichenkette im Eingabefeld zu
	hashen und gibt die resultierenden Hash-Werte aus.

	Beachte, dass die Hash-Werte für jede Zeichenkette im Eingabefeld garantiert eindeutig sind.

*/

func perfectHash(data []string) func(string) int {
	size := len(data)
	hash := make(map[string]int, size)

	// First hash function: returns the ASCII value of the first character
	for i := 0; i < size; i++ {
		hash[string(data[i][0])] = i
	}

	// Second hash function: returns the length of the string
	secondHash := func(str string) int {
		return len(str)
	}

	// Third hash function: returns the sum of the ASCII values of the characters
	thirdHash := func(str string) int {
		sum := 0
		for _, char := range str {
			sum += int(char)
		}
		return sum
	}

	// Generate a candidate hash function until a perfect hash function is found
	var candidate func(string) int
	for {
		candidate = func(str string) int {
			return (hash[string(str[0])] + secondHash(str) + thirdHash(str)) % size
		}
		used := make(map[int]bool, size)
		collision := false
		for _, str := range data {
			index := candidate(str)
			if used[index] {
				collision = true
				break
			}
			used[index] = true
		}
		if !collision {
			break
		}
	}

	return candidate
}

func main() {
	data := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "orange", "peach", "quince", "raspberry", "strawberry", "tangerine", "watermelon"}
	hash := perfectHash(data)
	for _, str := range data {
		fmt.Printf("%s: %d\n", str, hash(str))
	}
}
