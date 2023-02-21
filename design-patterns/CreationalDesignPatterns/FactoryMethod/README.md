# Factory Method

## Konzept

Das Factory-Entwurfsmuster, auch bekannt als FactoryMethod-Entwurfsmuster, ist eines der am weitesten verbreiteten und bekanntesten Entwurfsmuster in der Softwarebranche. Das Factory-Muster bietet eine Schnittstelle für die Erstellung von Objekten, erlaubt es aber Unterklassen zu bestimmen, welche Klasse instanziiert werden soll. Dies ermöglicht eine größere Flexibilität bei der Implementierung von Objekten, sowie eine einfachere Aktualisierung oder Herabstufung des zugrunde liegenden Typs, falls erforderlich.

## Zielsetzung

1. Delegieren der Erzeugung neuer Instanzen von Strukturen an einen anderen Teil des Programms
2. Arbeiten auf Schnittstellenebene statt mit konkreten Implementierungen
3. Gruppierung von Objektfamilien, um einen "family object creator" zu erhalten

Das Hauptziel des Factory-Musters besteht darin, die Erstellung neuer Instanzen von Strukturen an einen anderen Teil des Programms zu delegieren. Auf diese Weise wird der Code übersichtlicher und leichter zu pflegen. Außerdem arbeitet das Factory-Muster auf Schnittstellenebene und nicht mit konkreten Implementierungen, so dass es einfacher ist, die zugrunde liegende Leistung zu ändern, ohne den Client-Code zu beeinträchtigen.

### Pro

* Fördert die lose Kopplung zwischen Objekten, indem es den Prozess der Objekterstellung kapselt.
* Bietet eine klare Trennung der Verantwortlichkeiten, was die Wartung und Aktualisierung des Codes erleichtert.
* Unterstützt die Wiederverwendung von Code, da die Factory-Methode zur Erstellung von Objekten verwendet werden kann, die zur gleichen Objektfamilie gehören.
* Verbessert die Lesbarkeit und Organisation des Codes, da das Factory-Pattern die Trennung von Belangen fördert.
* Ermöglicht dynamische Objekterzeugung, da die Factory-Methode erweitert werden kann, um verschiedene Objekttypen auf der Grundlage der Eingabeparameter zu erzeugen.

### Contra

* Kann komplexer zu verstehen und zu implementieren sein als andere Entwurfsmuster.
* Fügt eine zusätzliche Abstraktionsebene hinzu, die einen Leistungs-Overhead verursachen und die Leistung des Systems verlangsamen kann.
* Dies kann zu einer engen Kopplung zwischen Objekten führen, wenn die Fabrikklasse eng mit den von ihr erstellten Objekten gekoppelt ist.
* Kann mehr Zeit und Aufwand für die Fehlersuche und -behebung erfordern, da das Fabrikmuster eine zusätzliche Komplexitätsschicht in den Code einführt.
* Kann übermäßig genutzt werden, was zu einer Aufblähung des Codes und unnötiger Abstraktion führt, was die Wartung und Aktualisierung des Codes erschweren kann.

Es ist wichtig, die Vor- und Nachteile der Verwendung des Factory Patterns sorgfältig abzuwägen und das Pattern zu wählen, das am besten zu den Anforderungen Ihres Systems passt. Das Factory-Pattern ist ein leistungsfähiges Werkzeug zur Vereinfachung der Erstellung von Objekten, aber es ist nicht unbedingt die beste Wahl für jede Situation.
