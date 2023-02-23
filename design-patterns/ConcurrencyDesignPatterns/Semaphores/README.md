# Semaphores

Eine Semaphore ist ein Synchronisationsmuster/Primitiv, das einer begrenzten Anzahl von Ressourcen einen gegenseitigen Ausschluss auferlegt.

Semaphore sind ein sehr allgemeiner Synchronisationsmechanismus, der verwendet werden kann, um Mutexe zu implementieren, den Zugriff auf mehrere Ressourcen zu beschränken, das Leser-Schreiber-Problem zu lösen, usw.

Es gibt keine Semaphore-Implementierung in Go's sync-Paket. Diese können leicht mit gepufferten Kanälen emuliert werden:

* Die Kapazität des gepufferten Kanals ist die Anzahl der Ressourcen, die synchronisiert werden sollen
* Die Länge (Anzahl der aktuell gespeicherten Elemente) des Kanals ist die Anzahl der Ressourcen, die aktuell verwendet werden
* Die Kapazität abzüglich der Länge des Kanals ist die Anzahl der freien Ressourcen (der ganzzahlige Wert der traditionellen Semaphore)
