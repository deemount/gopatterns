package main

import "log"

/*

	Beispiel:

	Diese Implementierung nimmt eine Textzeichenkette und eine Musterzeichenkette
	als Eingabe und gibt eine Scheibe von Ganzzahlen zurück, die die Indizes aller
	Vorkommen des Musters im Text darstellen. Wenn das Muster nicht gefunden wird,
	wird ein leeres Slice zurückgegeben.

	Die Implementierung verwendet das lps-Array, um den Suchprozess zu optimieren. Das
	lps-Array wird anhand der Musterzeichenkette berechnet und dient dazu, unnötige
	Vergleiche während der Suche zu vermeiden.

*/

func kmpSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)

	if m == 0 {
		return []int{}
	}

	// compute longest prefix suffix array
	lps := make([]int, m)
	i, j := 1, 0
	for i < m {
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	// perform search using lps array
	i, j = 0, 0
	indices := []int{}
	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}

		if j == m {
			indices = append(indices, i-m)
			j = lps[j-1]
		}
	}

	return indices
}

func main() {
	text := "Das ist ein sehr langer Text mit wenig inhalt, der dem zweck dient, ein Beispiel auszuführen, um Text zu durchsuchen. Ich schreibe dabei eigentlich nichts Genaues, aber soviel, dass es eben vielleicht reicht, ein brauchbares Ergebnis zu erhalten."
	result := kmpSearch(text, "ei")
	log.Printf("main: result %#v", result)
}
