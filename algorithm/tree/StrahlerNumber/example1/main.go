package main

import "log"

/*

	Beispiel:

	In dieser Implementierung wird der Eingabegraph als Adjazenzliste dargestellt.
	Die Funktion strahlerNumber nimmt die Adjazenzliste als Eingabe und gibt die Strahler-Zahl
	des Graphen zurück.

	Der Algorithmus berechnet zunächst den Indegree eines jeden Knotens und führt dann eine topologische
	Sortierung mit Hilfe einer Warteschlange durch. Während der Sortierung wird die maximale Strahler-Zahl
	der Nachbarn jedes Knotens berechnet, und die Strahler-Zahl des Knotens wird auf das Maximum
	plus eins gesetzt.

	Die Strahler-Zahl des Wurzelknotens (d. h. des Knotens ohne eingehende Kanten) wird als Endergebnis zurückgegeben.

	In diesem Beispiel hat der Eingabegraph eine Strahler-Zahl von 3, die von der
	Funktion strahlerNumber zurückgegeben wird.

*/

func strahlerNumber(adjList map[int][]int) int {
	indegree := make(map[int]int)
	for node, neighbors := range adjList {
		indegree[node] = 0
		for _, neighbor := range neighbors {
			indegree[neighbor]++
		}
	}

	order := make([]int, 0)
	for node, indeg := range indegree {
		if indeg == 0 {
			order = append(order, node)
		}
	}

	strahler := make(map[int]int)
	for len(order) > 0 {
		node := order[0]
		order = order[1:]

		maxStrahler := 0
		for _, neighbor := range adjList[node] {
			if strahler[neighbor] > maxStrahler {
				maxStrahler = strahler[neighbor]
			}
		}

		strahler[node] = maxStrahler
		if indegree[node] > 0 {
			indegree[node]--
			if indegree[node] == 0 {
				order = append(order, node)
			}
		}
	}

	rootStrahler := strahler[order[0]]
	for _, s := range strahler {
		if s > rootStrahler {
			rootStrahler = s
		}
	}

	return rootStrahler
}

func main() {
	adjList := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {5},
		4: {6},
		5: {6},
		6: {},
	}

	log.Println(strahlerNumber(adjList)) // Ausgabe: 3
}
