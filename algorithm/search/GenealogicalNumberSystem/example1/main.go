package main

import "log"

/*

	Beispiel:

	Dieser Algorithmus definiert eine Person-Struktur, die einen Namen, Vater, Mutter und
	das Feld ancestorNum enthält. Die setAncestorNum-Methode ordnet den direkten Vorfahren
	einer Person rekursiv Ahnennummern zu, wobei ein binäres Nummerierungssystem verwendet
	wird.

	Die Hauptfunktion erstellt einen Beispielstammbaum und ordnet den beiden Kindern Ahnennummern zu.
	Schließlich gibt sie die Ahnennummern für jedes Kind aus.

	In diesem Beispiel weist das Ahnennummerierungssystem dem ersten Kind (Tom) die Nummer 1 zu,
	die für den ersten Ahnen steht. Dem zweiten Kind (Sarah) wird die Nummer 2 zugewiesen,
	die für den zweiten Vorfahren steht. Dem Vater von Tom und Sarah (Mike) wird die Nummer 4
	zugewiesen, die für den vierten Vorfahren steht, und so weiter.

	Dieses System ermöglicht eine einfache Identifizierung der direkten Vorfahren einer Person
	und ihrer Beziehungen in einem Stammbaum.

*/

type Person struct {
	name           string
	father, mother *Person
	ancestorNum    int
}

func (p *Person) setAncestorNum(num int) {
	p.ancestorNum = num
	if p.father != nil {
		p.father.setAncestorNum(num * 2)
	}
	if p.mother != nil {
		p.mother.setAncestorNum(num*2 + 1)
	}
}

func main() {

	// create a sample family tree
	grandpa1 := &Person{name: "John"}
	grandma1 := &Person{name: "Jane"}
	grandpa2 := &Person{name: "Bob"}
	grandma2 := &Person{name: "Sue"}

	dad := &Person{name: "Mike", father: grandpa1, mother: grandma1}
	mom := &Person{name: "Lisa", father: grandpa2, mother: grandma2}

	child1 := &Person{name: "Tom", father: dad, mother: mom}
	child2 := &Person{name: "Sarah", father: dad, mother: mom}

	// assign ancestor numbers
	child1.setAncestorNum(1)
	child2.setAncestorNum(2)

	// print ancestor numbers
	log.Printf("%s's ancestor number is %d\n", child1.name, child1.ancestorNum)
	log.Printf("%s's ancestor number is %d\n", child2.name, child2.ancestorNum)
}
