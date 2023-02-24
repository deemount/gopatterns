package main

import "fmt"

/*

	Beispiel:

	In dieser Implementierung definiere ich eine Node-Struktur, die ein Schlüsselfeld sowie
	Zeiger auf linke und rechte Kindknoten enthält. Die Struktur BinaryTree enthält ein Feld
	Root, das ein Zeiger auf den Wurzelknoten des Baums ist.

	Die Methode Insert fügt dem Baum einen neuen Knoten mit dem angegebenen Schlüsselwert hinzu.
	Wenn der Baum leer ist, wird der neue Knoten der Wurzelknoten. Andernfalls wird die Insert-Methode
	rekursiv auf den entsprechenden Kindknoten aufgerufen, bis ein leerer Slot gefunden wird.

	Die Search-Methode gibt den Knoten mit dem angegebenen Schlüsselwert zurück. Wenn der Knoten gefunden
	wird, wird er zurückgegeben, andernfalls wird nil zurückgegeben. Die Suchmethode wird rekursiv auf
	dem entsprechenden untergeordneten Knoten aufgerufen, bis entweder der Knoten gefunden oder ein
	Blattknoten (leaf node) erreicht wird.

*/

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(key int) {
	if bt.Root == nil {
		bt.Root = &Node{Key: key}
	} else {
		bt.Root.insert(key)
	}
}

func (n *Node) insert(key int) {
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.insert(key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.insert(key)
		}
	}
}

func (bt *BinaryTree) Search(key int) *Node {
	return bt.Root.search(key)
}

func (n *Node) search(key int) *Node {
	if n == nil || n.Key == key {
		return n
	}

	if key < n.Key {
		return n.Left.search(key)
	}

	return n.Right.search(key)
}

func main() {
	// Create a new binary tree instance
	tree := &BinaryTree{}

	// Insert some nodes
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(9)

	// Search for a node
	node := tree.Search(7)
	if node != nil {
		fmt.Printf("Found node with key %d\n", node.Key)
	} else {
		fmt.Println("Node not found")
	}
}
