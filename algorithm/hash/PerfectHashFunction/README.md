# Perfect Hash Function

## Konzept

In der Informatik ist eine perfekte Hash-Funktion h für eine Menge S eine Hash-Funktion, die eindeutige Elemente in S auf eine Menge von m ganzen Zahlen abbildet, ohne dass es zu Kollisionen kommt. Mathematisch gesehen handelt es sich um eine injektive Funktion.

Perfekte Hash-Funktionen können verwendet werden, um eine Nachschlagetabelle mit konstanter Worst-Case-Zugriffszeit zu implementieren. Eine perfekte Hash-Funktion kann wie jede Hash-Funktion zur Implementierung von Hash-Tabellen verwendet werden, mit dem Vorteil, dass keine Kollisionsauflösung implementiert werden muss. Wenn es sich bei den Schlüsseln nicht um die Daten handelt und bekannt ist, dass die abgefragten Schlüssel gültig sind, müssen die Schlüssel nicht in der Nachschlagetabelle gespeichert werden, was Platz spart.

Nachteilig bei perfekten Hash-Funktionen ist, dass S für die Konstruktion der perfekten Hash-Funktion bekannt sein muss. Nicht-dynamische perfekte Hash-Funktionen müssen neu konstruiert werden, wenn sich S ändert. Für sich häufig ändernde S können dynamische perfekte Hash-Funktionen verwendet werden, was zusätzlichen Platz kostet. Der Platzbedarf zum Speichern der perfekten Hash-Funktion liegt bei O(n).

Die wichtigsten Leistungsparameter für perfekte Hash-Funktionen sind die Auswertungszeit, die konstant sein sollte, die Konstruktionszeit und die Darstellungsgröße.

## Anwendung

Eine perfekte Hash-Funktion, mit Werten in einem begrenzten Bereich, kann für effiziente Nachschlageoperationen verwendet werden, indem Schlüssel aus S (oder andere zugehörige Werte) in einer Nachschlagetabelle abgelegt werden, die durch die Ausgabe der Funktion indiziert ist. Man kann dann prüfen, ob ein Schlüssel in S vorhanden ist, oder einen mit diesem Schlüssel verknüpften Wert nachschlagen, indem man ihn in der entsprechenden Zelle der Tabelle sucht. Jeder dieser Nachschlagevorgänge benötigt im schlimmsten Fall konstante Zeit. Bei perfektem Hashing können die zugehörigen Daten mit einem einzigen Zugriff auf die Tabelle gelesen oder geschrieben werden.
