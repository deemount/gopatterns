# Builder

## Konzept

Builder ist ein Entwurfsmuster, das die Konstruktion eines komplexen Objekts von seiner Repräsentation trennt, so dass derselbe Konstruktionsprozess verschiedene Ausdrücke erzeugen kann. In Golang kann das Builder-Muster implementiert werden, indem man eine Builder-Struktur mit Methoden erstellt, um die Werte ihrer Felder zu setzen, und eine Build-Methode, die das zu erstellende Objekt zurückgibt.

Das Builder-Muster hilft uns dabei, komplexe Objekte zu konstruieren, ohne ihre Struktur direkt zu instanziieren oder die dafür erforderliche Logik zu schreiben. Es wird einen Algorithmus wiederverwenden, um viele Implementierungen einer Schnittstelle zu erstellen. Mit dem Builder-Muster können Sie verschiedene Typen und Darstellungen eines Objekts mit demselben Konstruktionscode erzeugen. Die Instanz des Builder-Musters liefert die öffnenden und schließenden geschweiften Klammern {} und verlässt die Instanz mit Nullwerten oder so komplex wie ein Objekt, das einige API-Aufrufe tätigen, Statistiken überprüfen und Objekte für seine Felder erstellen muss. Die BuildProcess Schnittstelle gibt an, was sie erfüllen muss, um Teil der möglichen Builder zu sein.

## Zielsetzung

* Abstrahieren komplexer Erstellungen, so dass die Objekterstellung vom Objektbenutzer getrennt ist.
* Schrittweises Erstellen eines Objekts durch Ausfüllen seiner Felder und Erstellen der eingebetteten Objekte.
* Wiederverwendung des Algorithmus zur Objekterstellung für viele Objekte.

### Pro

* Trennt die Konstruktion eines komplexen Objekts von seiner Darstellung, so dass derselbe Konstruktionsprozess verschiedene Darstellungen erzeugen kann.
* Fördert die Wiederverwendbarkeit des Codes, da der Builder zur Erstellung ähnlicher Objekte mit unterschiedlichen Konfigurationen verwendet werden kann.
* Verbessert die Lesbarkeit und Organisation des Codes, da das Builder-Pattern die Trennung von Belangen fördert.
* Es verbessert die Flexibilität des Codes, da neue Builder hinzugefügt werden können, um neue Objekttypen zu unterstützen, ohne den bestehenden Code zu ändern.
* Es vereinfacht die Objekterstellung, da der Builder die Details der Objektkonstruktion verbergen kann und die Erstellung komplexer Objekte erleichtert.

### Contra

* Führt zu erhöhter Komplexität, da die Anzahl der Builder und die Anzahl der konfigurierbaren Optionen schnell ansteigen kann.
* Erschweren das Verständnis des Codes, da das Builder-Muster mehrere Abstraktionsebenen einführen kann.
* Die Wartung und Aktualisierung des Codes wird erschwert, da Änderungen am Konstruktionsprozess oder an der Darstellung des Objekts Änderungen an mehreren Teilen des Codes erforderlich machen können.
* Der Code wird größer, da zusätzliche Klassen und Methoden für die Implementierung des Builder-Musters erforderlich sein können.

Es ist wichtig, die Vor- und Nachteile der Verwendung des Builder-Musters sorgfältig abzuwägen und das Muster zu wählen, das am besten zu den Anforderungen des Systems passt. Das Builder-Pattern ist ein leistungsfähiges Werkzeug zur Verwaltung der Erstellung komplexer Objekte, aber es ist möglicherweise nicht für jede Situation die beste Wahl.