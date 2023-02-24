# Least Frequency Used (LFU)

Cache-Evakuierungsalgorithmen werden häufig in Betriebssystemen, Datenbanken
und anderen Systemen verwendet, die Caches zur Beschleunigung der Ausführung, durch Zwischenspeichern von Daten, die von der Anwendung verwendet werden, nutzen.

Es gibt viele Verfahren, wie

* MRU (Most Recently Used),
* MFU (Most Frequently Used),
* LRU (Least Recently Used) und
* LFU (Least Frequently Used),

Die jeweils ihre Vor- und Nachteile haben und daher in bestimmten Szenarien verwendet werden.

Der bei weitem am häufigsten verwendete Algorithmus ist LRU, sowohl wegen seiner O(1)-Geschwindigkeit als auch wegen seiner großen Ähnlichkeit mit dem Verhalten, das die meisten
von den meisten Anwendungen erwartet wird.

Der LFU-Algorithmus hat auch ein Verhalten, das von vielen realen Workloads gewünscht wird. An vielen Stellen wird jedoch der LRU-Algorithmus dem LFU-Algorithmus vorgezogen, weil er eine geringere Laufzeitkomplexität von O(1) gegenüber O(log n).
