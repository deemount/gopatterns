# Quadtree

## Konzept

Ein Quadtree ist eine Baumdatenstruktur, bei der jeder interne Knoten genau vier Kinder hat. Quadtrees sind das zweidimensionale Gegenstück zu Octrees und werden meist verwendet, um einen zweidimensionalen Raum durch rekursive Unterteilung in vier Quadranten oder Regionen zu partitionieren. Die mit einer Blattzelle verbundenen Daten variieren je nach Anwendung, aber die Blattzelle stellt eine "Einheit interessanter räumlicher Informationen" dar.

Die unterteilten Regionen können quadratisch oder rechteckig sein, aber auch beliebige Formen haben. Diese Datenstruktur wurde 1974 von Raphael Finkel und J.L. Bentley als Quadtree bezeichnet. Eine ähnliche Aufteilung ist auch als Q-Tree bekannt. Alle Formen von Q-Trees haben einige gemeinsame Merkmale:

Sie unterteilen den Raum in anpassbare Zellen.
Jede Zelle (oder Bucket) hat eine maximale Kapazität. Wenn die maximale Kapazität erreicht ist, wird die Zelle geteilt.
Das Baumverzeichnis folgt der räumlichen Zerlegung des Quadtrees.
Eine Baumpyramide (T-Pyramide) ist ein "vollständiger" Baum; jeder Knoten der T-Pyramide hat vier untergeordnete Knoten, ausgenommen Blattknoten; alle Blätter befinden sich auf der gleichen Ebene, der Ebene, die den einzelnen Pixeln im Bild entspricht. Die Daten in einer Baumpyramide können kompakt in einem Array als implizite Datenstruktur gespeichert werden, ähnlich wie ein vollständiger Binärbaum kompakt in einem Array gespeichert werden kann.
