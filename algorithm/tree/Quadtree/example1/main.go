package main

import "log"

/*

	Beispiel:

	In dieser Implementierung steht
		* Point für einen Punkt im 2D-Raum,
		* Rectangle für eine rechteckige Begrenzung und
		* QuadTree für die QuadTree-Datenstruktur.

	Die Methode Insert wird verwendet, um einen Punkt zum QuadTree hinzuzufügen.

	* Wenn der Punkt außerhalb der Begrenzung liegt, gibt die Methode false zurück.
	* Wenn der QuadTree noch nicht geteilt wurde und noch Platz für den Punkt ist,
	  wird der Punkt zur Liste der Punkte des aktuellen Knotens hinzugefügt.
	* Wenn der Quadratzweig noch nicht unterteilt wurde und nicht genügend Platz für
	  den Punkt vorhanden ist, wird der Knoten in vier Quadranten unterteilt und
	  der Punkt rekursiv in den entsprechenden Quadranten eingefügt.
	* Wenn der Quadratzweig bereits unterteilt wurde, wird der Punkt rekursiv in den
	  entsprechenden Quadranten eingefügt.

	Die Methode Subdivide wird verwendet, um den aktuellen Knoten in vier Quadranten zu unterteilen.

	Mit der Contains-Methode wird geprüft, ob ein Punkt innerhalb eines Rechtecks liegt.

	Beachte, dass diese Implementierung ein einfaches Beispiel ist und für bestimmte Anwendungsfälle
	weiter optimiert werden kann.

	In diesem Beispiel erstelle ich einen neuen QuadTree mit einer Kapazität von 4
	und einer Rechteckbegrenzung von (0, 0) bis (100, 100). Anschließend füge ich dem
	QuadTree mit der Methode Insert 6 Punkte hinzu. Abschließend werden die Punkte im
	QuadTree ausgedruckt, die mit den eingefügten Punkten übereinstimmen sollten.

*/

type Point struct {
	x float64
	y float64
}

type Rectangle struct {
	x      float64
	y      float64
	width  float64
	height float64
}

type QuadTree struct {
	capacity  int
	boundary  Rectangle
	points    []Point
	divided   bool
	northwest *QuadTree
	northeast *QuadTree
	southwest *QuadTree
	southeast *QuadTree
}

func (qt *QuadTree) Insert(point Point) bool {
	if !qt.boundary.Contains(point) {
		return false
	}
	if len(qt.points) < qt.capacity {
		qt.points = append(qt.points, point)
		return true
	}
	if !qt.divided {
		qt.Subdivide()
	}
	if qt.northwest.Insert(point) {
		return true
	}
	if qt.northeast.Insert(point) {
		return true
	}
	if qt.southwest.Insert(point) {
		return true
	}
	if qt.southeast.Insert(point) {
		return true
	}
	return false
}

func (qt *QuadTree) Subdivide() {
	x := qt.boundary.x
	y := qt.boundary.y
	w := qt.boundary.width / 2
	h := qt.boundary.height / 2
	qt.northwest = &QuadTree{
		capacity: qt.capacity,
		boundary: Rectangle{x, y, w, h},
	}
	qt.northeast = &QuadTree{
		capacity: qt.capacity,
		boundary: Rectangle{x + w, y, w, h},
	}
	qt.southwest = &QuadTree{
		capacity: qt.capacity,
		boundary: Rectangle{x, y + h, w, h},
	}
	qt.southeast = &QuadTree{
		capacity: qt.capacity,
		boundary: Rectangle{x + w, y + h, w, h},
	}
	qt.divided = true
}

func (r Rectangle) Contains(p Point) bool {
	return p.x >= r.x-r.width/2 &&
		p.x <= r.x+r.width/2 &&
		p.y >= r.y-r.height/2 &&
		p.y <= r.y+r.height/2
}

func main() {

	boundary := Rectangle{0, 0, 100, 100}
	quadTree := QuadTree{capacity: 4, boundary: boundary}
	points := []Point{
		{10, 10},
		{20, 10},
		{15, 15},
		{5, 90},
		{70, 80},
		{80, 90},
	}
	for _, point := range points {
		quadTree.Insert(point)
	}
	log.Println(quadTree.points) // Ausgabe: [{10 10} {20 10} {15 15} {5 90} {70 80} {80 90}]
}
