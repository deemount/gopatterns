# Balking

Ist ein Entwurfsmuster, das eine Aktion für ein Objekt nur dann ausführt, wenn sich das Objekt in einem bestimmten Zustand befindet.

Wenn ein Objekt beispielsweise ZIP-Dateien liest und eine aufrufende Methode eine Get-Methode für das Objekt aufruft, obwohl die ZIP-Datei nicht geöffnet ist, würde das Objekt die Anforderung "blockieren".

Balking wird eher als Anti-Muster, denn als Entwurfsmuster betrachtet. Wenn ein Objekt seine API nicht unterstützen kann, sollte es entweder die API einschränken, so dass der beanstandete Aufruf nicht verfügbar ist, oder so, dass der Aufruf ohne Einschränkung erfolgen kann. 

Es sollte:

* in einem "vernünftigen Zustand" erstellt werden
* sich selbst nicht verfügbar machen, bis sie sich in einem "gesunden" Zustand befindet
* eine sog. ***Facade*** werden und einem Objekt antworten, das sich in einem "gesunden" Zustand befindet

## Zusammenfassung

Dieses Entwurfsmuster verhindert unnötige Arbeit, indem es Operationen überspringt, die den Objektzustand nicht verändern würden, und verbessert so die Effizienz und Leistung.
