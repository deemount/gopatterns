package main

import (
	"log"
	"math"
)

/*

	Beispiel:

	Diese Implementierung erstellt einen 2D-k-d-Baum aus einer Scheibe von
	Punkten und sucht dann nach dem nächstgelegenen Punkt zu einem Zielpunkt
	unter Verwendung des k-d-Baums. Die Point-Struktur stellt einen Punkt im 2D-Raum dar
	und die KDTree-Struktur stellt einen k-d-Baum dar.

	Die Funktion NewKDTree konstruiert den k-d-Baum durch rekursive Partitionierung der
	Punkte entlang alternierender Dimensionen. Die NearestNeighbor-Methode sucht den k-d-Baum
	nach dem nächstgelegenen Punkt zu einem Zielpunkt mit Hilfe der nearestNeighbor-Hilfsfunktion,
	die den k-d-Baum rekursiv durchläuft und bei Bedarf Teilbäume ausschneidet.

	Die Methode getAxisDistance berechnet den Abstand zwischen einem Punkt und einem anderen Punkt
	entlang einer bestimmten Achse.

*/

type Point struct {
	X, Y float64
}

type KDTree struct {
	root *node
}

type node struct {
	point Point
	left  *node
	right *node
	axis  int
}

func NewKDTree(points []Point) *KDTree {
	if len(points) == 0 {
		return nil
	}

	return &KDTree{root: buildSubtree(points, 0)}
}

func buildSubtree(points []Point, depth int) *node {
	if len(points) == 0 {
		return nil
	}

	axis := depth % 2
	medianIdx := len(points) / 2

	median := points[medianIdx]
	left := buildSubtree(points[:medianIdx], depth+1)
	right := buildSubtree(points[medianIdx+1:], depth+1)

	return &node{point: median, left: left, right: right, axis: axis}
}

func (tree *KDTree) NearestNeighbor(target Point) Point {
	return nearestNeighbor(tree.root, target, tree.root.point, math.MaxFloat64)
}

func nearestNeighbor(node *node, target Point, bestPoint Point, bestDistance float64) Point {
	if node == nil {
		return bestPoint
	}

	distance := node.point.distance(target)
	if distance < bestDistance {
		bestDistance = distance
		bestPoint = node.point
	}

	axisDistance := target.getAxisDistance(node.axis, node.point)
	if axisDistance < bestDistance {
		bestPoint = nearestNeighbor(node.left, target, bestPoint, bestDistance)
		bestDistance = bestPoint.distance(target)

		if target.getAxisDistance(node.axis, node.point) < bestDistance {
			bestPoint = nearestNeighbor(node.right, target, bestPoint, bestDistance)
		}
	}

	return bestPoint
}

func (point Point) distance(other Point) float64 {
	dx := point.X - other.X
	dy := point.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (point Point) getAxisDistance(axis int, other Point) float64 {
	if axis == 0 {
		return point.X - other.X
	}
	return point.Y - other.Y
}

func main() {
	// Define a slice of points to build the k-d tree
	points := []Point{
		{X: 2, Y: 3},
		{X: 5, Y: 4},
		{X: 9, Y: 6},
		{X: 4, Y: 7},
		{X: 8, Y: 1},
		{X: 7, Y: 2},
	}

	// Build the k-d tree
	kdTree := NewKDTree(points)

	// Define a target point
	target := Point{X: 6, Y: 5}

	// Find the nearest neighbor to the target point
	nearestNeighbor := kdTree.NearestNeighbor(target)

	// Print the nearest neighbor
	log.Printf("The nearest neighbor to (%f, %f) is (%f, %f)\n", target.X, target.Y, nearestNeighbor.X, nearestNeighbor.Y)
}
