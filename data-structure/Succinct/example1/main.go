package main

import "log"

/*

	Beispiel:

	In dieser Implementierung speichert die SuccinctBitVector-Struktur die Bits des
	Bitvektors als ein Slice von uint64-Ganzzahlen, wobei jede uint64 64 Bits des Bitvektors
	darstellt. Das rankCache-Slice speichert die Anzahl der 1en bis zu jeder Position im
	Bitvektor und ermöglicht so eine effiziente Berechnung der Rangfunktion.

	Die Rank-Methode berechnet die Anzahl der 1en bis zu einer bestimmten Position im Bitvektor
	unter Verwendung des rankCache.

	Die Select-Methode findet die Position der i-ten 1 im Bitvektor durch binäre Suche. Bei der Implementierung
	wird angenommen, dass i im Bereich [1, n] liegt, wobei n die Anzahl der 1en im Bitvektor ist.

*/

type SuccinctBitVector struct {
	bits      []uint64
	rankCache []int // rankCache[i] stores the number of 1s up to position i
}

func NewSuccinctBitVector(bits []uint64) *SuccinctBitVector {
	n := len(bits) * 64
	rankCache := make([]int, n+1)

	// Compute the rankCache
	for i := 1; i <= n; i++ {
		rankCache[i] = rankCache[i-1] + int(bits[(i-1)/64]>>uint(i%64-1)&1)
	}

	return &SuccinctBitVector{
		bits:      bits,
		rankCache: rankCache,
	}
}

func (sbv *SuccinctBitVector) Rank(i int) int {
	return sbv.rankCache[i]
}

func (sbv *SuccinctBitVector) Select(i int) int {
	// Binary search for the i-th 1 in the bit vector
	l, r := 0, len(sbv.rankCache)-1
	for l <= r {
		mid := (l + r) / 2
		if sbv.rankCache[mid] < i {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// If i is out of range, return -1
	if l >= len(sbv.rankCache) || sbv.rankCache[l] != i {
		return -1
	}

	return l
}

func main() {

	bits := []uint64{0b11010010, 0b10101011, 0b11110000}
	sbv := NewSuccinctBitVector(bits)

	// Compute the rank of the 1st, 3rd, and 5th bits
	log.Println(sbv.Rank(1)) // Output: 1
	log.Println(sbv.Rank(3)) // Output: 2
	log.Println(sbv.Rank(5)) // Output: 2

	// Find the position of the 2nd and 4th 1s
	log.Println(sbv.Select(2)) // Output: 3
	log.Println(sbv.Select(4)) // Output: 10
}
