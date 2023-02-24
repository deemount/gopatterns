package main

import "log"

/*

	Beispiel:

	Dieser Algorithmus nimmt ein Array von Zeichenketten als Eingabe und gibt das am
	häufigsten vorkommende Element im Array zurück. Er erstellt zunächst eine Karte,
	um die Häufigkeit jedes Elements in der Matrix zu speichern und durchläuft dann die Matrix,
	um die Häufigkeit jedes Elements zu zählen.

	Schließlich wird die Map durchlaufen, um das Element mit der höchsten Häufigkeit zu finden
	und zurückzugegeben.
*/

func MostFrequent(arr []string) string {
	freqMap := make(map[string]int)
	maxFreq := 0
	mostFreq := ""

	// Count the frequency of each element in the array
	for _, elem := range arr {
		freqMap[elem]++
	}

	// Find the element with the highest frequency
	for elem, freq := range freqMap {
		if freq > maxFreq {
			maxFreq = freq
			mostFreq = elem
		}
	}

	return mostFreq
}

func main() {

	most := MostFrequent([]string{"hello", "world", "foo", "bar", "hot", "summer", "hello", "world", "planet", "earth", "Hello"})
	log.Printf("main: most %s", most)

}
