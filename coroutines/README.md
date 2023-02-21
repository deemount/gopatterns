# Co Routines

Benutzer von C#, Lua, Python und anderen Programmiersprachen bemerken vielleicht die Ähnlichkeit zwischen den Generatoren/Iteratoren von Go und den Coroutines.

Die Unterschiede zwischen Coroutinen und Goroutinen sind:

* Goroutinen implizieren Parallelität
* Coroutinen tun dies im Allgemeinen nicht
* Goroutinen kommunizieren über Kanäle
* Coroutinen kommunizieren über Yield- und Resume-Operationen

Im Allgemeinen sind Goroutinen viel leistungsfähiger als Coroutinen.  Insbesondere ist es einfach, Coroutine-Logik auf Goroutinen zu portieren und dabei Parallelität zu erreichen
