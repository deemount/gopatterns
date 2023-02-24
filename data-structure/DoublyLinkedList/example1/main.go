package main

import (
	"log"
)

/*

	Beispiel:

	In dieser Implementierung stellt die Node-Struktur einen einzelnen Knoten in der doppelt verketteten
	Liste dar und enthält einen Verweis auf seine Daten sowie Zeiger auf den vorherigen und den nächsten Knoten.
	Die DoublyLinkedList-Struktur repräsentiert die Liste selbst, mit Zeigern auf ihre Kopf- und Endknoten.

	Die Methoden PushBack und PushFront fügen einen neuen Knoten am Ende bzw. am Anfang der Liste hinzu. Die Methoden
	PopBack und PopFront entfernen die Daten vom Ende bzw. Anfang der Liste und geben sie zurück.

	Beachte, dass die Liste bei dieser Implementierung jeden Datentyp enthalten kann, indem ein interface{}-Feld in der
	Node-Struktur verwendet wird. Dies bedeutet jedoch auch, dass ich die zurückgegebenen Daten auf den
	entsprechenden Typ umwandeln muss, wenn ich sie verwende.

*/

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) PushBack(data interface{}) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	node.prev = list.tail
	list.tail = node
}

func (list *DoublyLinkedList) PushFront(data interface{}) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.head.prev = node
	node.next = list.head
	list.head = node
}

func (list *DoublyLinkedList) PopBack() interface{} {
	if list.tail == nil {
		return nil
	}

	data := list.tail.data

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}

	return data
}

func (list *DoublyLinkedList) PopFront() interface{} {
	if list.head == nil {
		return nil
	}

	data := list.head.data

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.next
		list.head.prev = nil
	}

	return data
}

func main() {
	// Create a new doubly linked list
	list := DoublyLinkedList{}

	// Add some nodes to the list
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	// Add a node to the beginning of the list
	list.PushFront(0)

	// Remove a node from the end of the list
	last := list.PopBack()

	log.Println("Removed node:", last)

	// Remove a node from the beginning of the list
	first := list.PopFront()

	log.Println("Removed node:", first)

	// Print the remaining nodes in the list
	node := list.head
	for node != nil {
		log.Println("Node data:", node.data)
		node = node.next
	}
}
