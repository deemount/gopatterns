# Knuth Morris Pratt

## Konzept

In der Informatik sucht der Knuth-Morris-Pratt-Algorithmus (oder KMP-Algorithmus) nach dem Vorkommen eines "Wortes" W in einer Haupt-"Textkette" S, indem er die Beobachtung nutzt, dass bei einer Nichtübereinstimmung das Wort selbst genügend Informationen enthält, um zu bestimmen, wo die nächste Übereinstimmung beginnen könnte, so dass eine erneute Prüfung der zuvor übereinstimmenden Zeichen umgangen wird.

Der Algorithmus wurde von James H. Morris erdacht und "einige Wochen später" von Donald Knuth unabhängig von der Automatentheorie entdeckt. Morris und Vaughan Pratt veröffentlichten 1970 einen technischen Bericht. Die drei veröffentlichten den Algorithmus 1977 auch gemeinsam. Unabhängig davon entdeckte Matijasewitsch 1969 einen ähnlichen Algorithmus, der von einer zweidimensionalen Turing-Maschine kodiert wurde, während er ein Problem der Erkennung von Zeichenkettenmustern in einem binären Alphabet untersuchte. Dies war der erste Algorithmus mit linearer Zeit für String-Matching.
