package main

import "fmt"

/*

	Beispiel:

	In diesem Beispiel definiere ich eine Node-Struktur mit einem Wert und
	einem Zeiger auf den nächsten Knoten in der Liste. Außerdem definiere
	ich eine LinkedList-Struktur mit einem Zeiger auf den Kopf der Liste.

	Die Insert-Methode fügt einen neuen Knoten an das Ende der Liste an.
	Wenn die Liste leer ist, setzt sie den neuen Knoten an die Spitze. Andernfalls
	durchläuft sie die Liste, um den letzten Knoten zu finden und setzt ihren
	nächsten Zeiger auf den neuen Knoten.

	Die Methode printList durchläuft die Liste und gibt den Wert jedes Knotens aus.

	In der Hauptfunktion erstellen wir eine neue verknüpfte Liste, fügen drei Werte
	ein und drucken die Liste aus. Die Ausgabe sollte wie folgt aussehen:

*/

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insert(value int) {
	newNode := &Node{value: value}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}
	ll.insert(1)
	ll.insert(2)
	ll.insert(3)
	ll.printList()
}
