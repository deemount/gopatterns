# Interpreter

Das Interpreter-Muster legt fest, wie benutzerdefinierte Sprachen auszuwerten sind. Die Beispiele sind grenzenlos, zum Beispiel SQL.

## Beispiele

* Spezialisierte Datenbankabfragesprachen wie SQL
* Spezialisierte Computersprachen, die häufig zur Beschreibung von Kommunikationsprotokollen verwendet werden
* Die meisten Allzweck-Computersprachen enthalten mehrere spezialisierte Sprachen

### Welche Probleme kann das Interpreter-Entwurfsmuster lösen?

Es soll eine Grammatik für eine einfache Sprache definiert werden damit Sätze in der Sprache interpretiert werden können.
Wenn ein Problem sehr häufig auftritt, könnte man überlegen, es als Satz in einer einfachen Sprache (Domain Specific Languages) darzustellen, so dass ein Interpreter das Problem durch Interpretation des Satzes lösen kann.

Zum Beispiel, wenn viele verschiedene oder komplexe Suchausdrücke angegeben werden müssen. Sie direkt in eine Klasse zu implementieren (fest zu verdrahten) ist unflexibel, weil die Klasse dadurch an bestimmte Ausdrücke gebunden ist und es unmöglich ist, neue Ausdrücke zu spezifizieren oder bestehende unabhängig von der Klasse zu ändern (ohne sie ändern zu müssen).

### Welche Lösung beschreibt das Interpreter-Entwurfsmuster?

* Definieren einer Grammatik für eine einfache Sprache, indem eine Expression-Klassenhierarchie definiert und eine interpret()-Operation implementiert wird
* Repräsentiert einen Satz in der Sprache durch einen abstrakten Syntaxbaum (AST), der aus Expression-Instanzen besteht
* Interpretiert einen Satz, indem interpret() auf dem AST aufgerufen wird
* Die Expression-Objekte werden rekursiv zu einer Composite/Tree-Struktur zusammengesetzt, die als abstrakter Syntaxbaum bezeichnet wird (siehe Composite-Entwurfsmuster)
* Das Interpreter-Muster beschreibt nicht, wie man einen abstrakten Syntaxbaum aufbaut. Dies kann entweder manuell durch einen Client oder automatisch durch einen Parser geschehen
