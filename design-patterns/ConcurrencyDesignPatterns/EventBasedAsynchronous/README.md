# Event Based Asynchronous

Das ereignisbasierte asynchrone Muster stellt die Vorteile von Multithreading-Anwendungen zur Verfügung und verbirgt gleichzeitig viele der komplexen Probleme, die mit einem Multithreading-Design verbunden sind. Die Verwendung einer Klasse, die dieses Muster unterstützt, kann es mir ermöglichen,:

* Zeitaufwändige Aufgaben wie Downloads und Datenbankoperationen "im Hintergrund" auszuführen, ohne die Anwendung zu unterbrechen.
* Mehrere Operationen gleichzeitig auszuführen und Benachrichtigungen zu erhalten, wenn eine Operation abgeschlossen ist.
* Warten auf die Verfügbarkeit von Ressourcen, ohne die Anwendung anzuhalten ("aufzuhängen").
* Kommuniziere mit ausstehenden asynchronen Operationen über das bekannte Ereignis- und Delegatenmodell.

## Anwendbarkeit

Verwende das/die ereignisbasierte(n) asynchrone(n) Muster, wenn zeitaufwändige Aufgaben im Hintergrund ausgeführt werden müssen, ohne die laufende Anwendung zu unterbrechen.
