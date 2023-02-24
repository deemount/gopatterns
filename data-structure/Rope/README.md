# Datenstrukturen

## Rope

In der Computerprogrammierung ist ein Seil oder eine Schnur eine Datenstruktur, die aus kleineren Zeichenketten besteht und dazu dient, eine sehr lange Zeichenkette effizient zu speichern und zu bearbeiten. Ein Textbearbeitungsprogramm kann beispielsweise ein Seil verwenden, um den zu bearbeitenden Text darzustellen, so dass Vorgänge wie Einfügen, Löschen und zufälliger Zugriff effizient durchgeführt werden können.

Ein Seil (Rope) ist eine Art binärer Baum, bei dem jedes Blatt (Endknoten) eine Zeichenkette und eine Länge (auch "Gewicht" genannt) enthält, und jeder weiter oben im Baum liegende Knoten enthält die Summe der Längen aller Blätter in seinem linken Teilbaum. Ein Knoten mit zwei Kindern teilt also die gesamte Zeichenkette in zwei Teile: Der linke Teilbaum speichert den ersten Teil der Zeichenkette, der rechte Teilbaum den zweiten Teil der Zeichenkette, und das Gewicht eines Knotens ist die Länge des ersten Teils.

Bei Seiloperationen wird davon ausgegangen, dass die in den Knoten gespeicherten Zeichenketten im typischen nicht-destruktiven Fall konstante, unveränderliche Objekte sind, die ein gewisses Copy-on-Write-Verhalten ermöglichen. Blattknoten werden in der Regel als einfache Strings mit fester Länge implementiert, die mit einem Referenzzähler versehen sind, um sie zu löschen, wenn sie nicht mehr benötigt werden, obwohl auch andere Garbage-Collection-Methoden verwendet werden können.
