# Double Hashing

## Konzept

Double Hashing ist eine Computerprogrammierungstechnik, die in Verbindung mit offener Adressierung in Hash-Tabellen verwendet wird, um Hash-Kollisionen aufzulösen, indem ein sekundärer Hash des Schlüssels als Offset verwendet wird, wenn eine Kollision auftritt. Double Hashing mit offener Adressierung ist eine klassische Datenstruktur in einer Tabelle T.

Das doppelte Hash-Verfahren verwendet einen Hash-Wert als Index in der Tabelle und schreitet dann wiederholt ein Intervall vorwärts, bis der gewünschte Wert gefunden, eine leere Stelle erreicht oder die gesamte Tabelle durchsucht worden ist; dieses Intervall wird jedoch durch eine zweite, unabhängige Hash-Funktion festgelegt. Im Gegensatz zu den alternativen Kollisionsauflösungsmethoden des linearen Sondierens und des quadratischen Sondierens, hängt das Intervall von den Daten ab, so dass Werte, die auf dieselbe Stelle abgebildet werden, unterschiedliche Bucket-Sequenzen haben; dadurch werden wiederholte Kollisionen und die Auswirkungen der Clusterbildung minimiert.
