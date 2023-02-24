package main

import "log"

/*

	Beispiel:

	In diesem Beispiel definiere ich eine Rope-Struktur, die einen linken und
	einen rechten Kindknoten hat (die ebenfalls Rope-Strukturen sind).
	Ein Datenfeld für die eigentlichen Zeichenkettendaten (nur in Blattknoten)
	und ein len-Feld, um die Gesamtlänge des Seils zu verfolgen.

	Dann definiere ich Methoden, um ein Zeichen an einem bestimmten Index zu
	erhalten (Get), zwei Strings miteinander zu verketten (Concat) und einen
	Teilstring des Strings zu erhalten (Substr).

	In der Hauptfunktion erstelle ich zwei Seile mit den Zeichenketten "hello" und "world",
	verketten sie miteinander und teste dann die Methoden Substr und Get an dem
	resultierenden Seil.

*/

type Rope struct {
	left  *Rope
	right *Rope
	data  string
	len   int
}

func NewRope(data string) *Rope {
	return &Rope{data: data, len: len(data)}
}

func (r *Rope) Get(i int) byte {
	if r.left == nil && r.right == nil {
		return r.data[i]
	} else if i < r.left.len {
		return r.left.Get(i)
	} else {
		return r.right.Get(i - r.left.len)
	}
}

func (r *Rope) Concat(other *Rope) *Rope {
	return &Rope{
		left:  r,
		right: other,
		len:   r.len + other.len,
	}
}

func (r *Rope) Substr(start, end int) string {
	if r.left == nil && r.right == nil {
		return r.data[start:end]
	} else if end <= r.left.len {
		return r.left.Substr(start, end)
	} else if start >= r.left.len {
		return r.right.Substr(start-r.left.len, end-r.left.len)
	} else {
		return r.left.Substr(start, r.left.len) + r.right.Substr(0, end-r.left.len)
	}
}

func main() {
	r1 := NewRope("hello")
	r2 := NewRope("world")
	rope := r1.Concat(r2)

	log.Println(rope.Substr(1, 8))   // Ausgabe: elloworl
	log.Println(string(rope.Get(4))) // Ausgabe: o
}
