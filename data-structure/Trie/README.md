# Trie

Ein Trie oder Präfixbaum ist eine Datenstruktur, die in der Informatik zum Suchen nach Zeichenketten verwendet wird. Es handelt sich dabei um einen speziellen Suchbaum zur gleichzeitigen Speicherung mehrerer Zeichenketten. Dabei beinhalten Tries eine Art der Datenkompression, da gemeinsame Präfixe der Zeichenketten nur einmal gespeichert werden.

Ein Trie wird über eine Menge von beliebigen Zeichenketten aufgebaut. Jede ausgehende Kante eines Knotens innerhalb eines Tries ist mit einem einzelnen Zeichen versehen, sodass ein Pfad beginnend bei der Wurzel bis zu einem Blatt im Trie eine der Zeichenketten darstellt, aus denen der Baum konstruiert worden ist.

Tries finden ihre Anwendung im Bereich des Information Retrieval. Dort werden sie zur Indexierung von Texten verwendet, um effizient bestimmte Anfragen an den Text zu beantworten.

## Information Retrieval

betrifft das Wiederauffinden von Information, meist durch Abruf aus Datenbanken. Das Fachgebiet beschäftigt sich mit computergestütztem Suchen nach komplexen Inhalten (also nicht z. B. nach Einzelwörtern) und fällt in die Bereiche Informationswissenschaft, Informatik und Computerlinguistik

Komplexe Texte oder Bilddaten, die in großen Datenbanken gespeichert werden, sind für Außenstehende zunächst nicht zugänglich oder abrufbar. Das Wort retrieval bedeutet auf Deutsch Abruf bzw. Wiederauffinden. Beim IR geht es also darum, bestehende Informationen wieder aufzufinden. Etwas anderes wäre das Entdecken neuer Strukturen: Das gehört zur Knowledge Discovery in Databases mit Data-Mining und Text Mining.

## Anwendungsbeispiele

Mit Hilfe von Tries können unterschiedliche Anfragen an eine gegebene Menge verschiedener Zeichenketten  gestellt werden.

Beispielhafte Anfragen könnten sein:

Existenzanfragen der Art „Enthält S das Muster M?“
Präfixanfragen der Art „Welche Zeichenketten in S beginnen mit dem Muster M?“
Nachfolger- und Vorgängeranfragen wie „Welche Zeichenketten in S sind lexikographische Nachfolger bzw. Vorgänger des Musters M?

Eine mögliche Verwendung von Tries könnte beispielsweise die Realisierung von Suchanfragen innerhalb einer Kontakte- bzw. Telefonbuch-App für Smartphones sein.

Mit Hilfe des Tries kann eine Personensuche nach Namen erfolgen (Existenzanfragen). Ebenso können bei Eingabe eines Namens bereits Kontakte angezeigt werden, deren Namen mit den bisher eingegebenen Buchstaben beginnen (Präfixanfragen). Des Weiteren können Kontakte gefunden werden, die im Telefonbuch hinter bzw. vor der angefragten Person stehen (Nachfolger- und Vorgängeranfragen).
