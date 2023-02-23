# Strukturelle Entwurfsmuster

Strukturelle Entwurfsmuster bilden in der Softwareentwicklung Entwurfsmuster, die durch das Herstellen von Beziehungen zwischen Entitäten, den Entwurf der Software erleichtern.

## Kurzbeschreibung

* Adapter: adaptiert eine Schnittstelle für eine Klasse in eine andere, die der Client erwartet:
  * Adapter-Pipeline: Verwendet mehrere Adapter, um so den Code auszutesten.
  * Nachrüstungsschnittstellenmuster (Retrofit Interface Pattern): Besteht aus einem Adapter, der gleichzeitig als eine neue Schnittstelle für mehrere Klassen verwendet wird.
* Aggregat (Aggregat): Ist eine Version des Kompositum, die zusätzlich Methoden zur Gruppierung von abgeleiteten Klassen, sprich Nachkommen zur Verfügung stellt
* Die Brücke (Bridge): Entkoppelt eine Abstraktion und deren Implementation, sodass beide sich unabhängig voneinander weiterentwickeln können
* Dekorierer (Dekorierer): Ermöglicht der Klasse zusätzliche Funktionalität während der Laufzeit hinzuzufügen, wobei durch Ableiten die Klassenanzahl exponentiell ansteigt.
* Fassade (Facade): Erzeugt eine vereinfachte Schnittstelle einer anderen Schnittstelle, um die Benutzung letzterer zu vereinfachen.
* Fliegengewicht (Flyweight): Bei diesem Muster teilen sich eine große Anzahl von Objekten ein Objekt der allgemeinen Eigenschaften, um Sicherungsspeicher einzusparen
* Kompositum (Composite): Eine Baumstruktur von Objekten, in der jedes Objekt dieselbe Schnittstelle verwendet.
* Stellvertreter (Proxy): Hier fungiert eine Klasse als eine Schnittstelle für etwas anderes.

Folgende strukturelle Entwurfsmuster für Go sind mit Beispielen enthalten

* Adapter
* Aggregate
* Bridge
* Composite
* Decorator
* Facade
* Flyweight
* Proxy
