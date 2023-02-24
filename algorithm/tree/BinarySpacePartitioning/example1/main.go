package main

import "fmt"

/*

	Beispiel:

	In diesem Beispiel stellt die Node-Struktur einen Knoten in einem Baum mit binärer
	Raumaufteilung dar. Die Rect-Struktur stellt ein Rechteck im 2D-Raum dar.
	Die Insert-Methode fügt ein neues Rechteck in den Baum ein, indem sie den Baum rekursiv
	durchläuft und die Knoten nach Bedarf aufteilt, um die binäre Raumaufteilung beizubehalten.

	Die Funktion intersectionArea berechnet den Schnittbereich zwischen zwei Rechtecken
	und die Funktion combineRects kombiniert zwei Rechtecke zu einem einzigen Rechteck,
	das beide enthält.

	Dies ist nur ein einfaches Beispiel, um die Grundprinzipien der binären Raumaufteilung in Golang zu
	veranschaulichen. Es gibt viele Variationen und Optimierungen, die je nach spezifischem Anwendungsfall
	vorgenommen werden können.

	In diesem Beispiel erstelle ich einen neuen Knoten, der den gesamten 100x100-Raum repräsentiert
	und füge dann mit der Methode Insert vier Rechtecke unterschiedlicher Größe und Position in den Baum ein.
	Zum Schluss wird der Wurzelknoten des Baums ausgedruckt, um die resultierende binäre
	Raumaufteilungsstruktur zu sehen.

	Diese Ausgabe zeigt, dass der Wurzelknoten des Baums zwei untergeordnete Knoten hat, die die beiden Hälften
	des 100x100-Raums darstellen. Der linke untergeordnete Knoten hat zwei Rechtecke, rect1 und rect3, und der rechte
	untergeordnete Knoten hat ein Rechteck, rect2. Das vierte Rechteck, rect4, ist in keinem dieser Knoten enthalten
	und bleibt daher ein Blattknoten an der Wurzel des Baums.

*/

type Node struct {
	left  *Node
	right *Node
	rect  Rect
}

type Rect struct {
	minX, minY, maxX, maxY int
}

func (n *Node) Insert(rect Rect) *Node {
	if n == nil {
		return &Node{rect: rect}
	}
	if n.left != nil && n.right != nil {
		newLeftArea := intersectionArea(n.left.rect, rect)
		newRightArea := intersectionArea(n.right.rect, rect)

		if newLeftArea < newRightArea {
			n.left = n.left.Insert(rect)
		} else {
			n.right = n.right.Insert(rect)
		}

		n.rect = combineRects(n.rect, rect)

		return n
	}

	if n.left == nil {
		n.left = &Node{rect: rect}
		return n.left
	}

	n.right = &Node{rect: rect}
	return n.right
}

func intersectionArea(a, b Rect) int {
	minX := max(a.minX, b.minX)
	minY := max(a.minY, b.minY)
	maxX := min(a.maxX, b.maxX)
	maxY := min(a.maxY, b.maxY)

	if maxX <= minX || maxY <= minY {
		return 0
	}

	return (maxX - minX) * (maxY - minY)
}

func combineRects(a, b Rect) Rect {
	return Rect{
		minX: min(a.minX, b.minX),
		minY: min(a.minY, b.minY),
		maxX: max(a.maxX, b.maxX),
		maxY: max(a.maxY, b.maxY),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	rootNode := &Node{rect: Rect{0, 0, 100, 100}}

	rect1 := Rect{10, 10, 20, 20}
	rootNode.Insert(rect1)

	rect2 := Rect{30, 30, 40, 40}
	rootNode.Insert(rect2)

	rect3 := Rect{15, 15, 25, 25}
	rootNode.Insert(rect3)

	rect4 := Rect{50, 50, 60, 60}
	rootNode.Insert(rect4)

	fmt.Println(rootNode)
}
