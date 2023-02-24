# Cuckoo Hashing

## Konzept

Cuckoo Hashing ist ein Schema in der Computerprogrammierung zur Auflösung von Hash-Kollisionen von Werten von Hash-Funktionen in einer Tabelle, mit im schlimmsten Fall konstanter Suchzeit. Der Name leitet sich vom Verhalten einiger Kuckucksarten ab, bei denen das Kuckucksküken die anderen Eier oder Jungtiere aus dem Nest stößt, wenn es schlüpft - eine Variante des Verhaltens, die als Brutparasitismus bezeichnet wird.

## Ausführung

Cuckoo-Hashing ist eine Form der offenen Adressierung, bei der jede nicht leere Zelle einer Hashtabelle einen Schlüssel oder ein Schlüssel-Wert-Paar enthält. Eine Hash-Funktion wird verwendet, um den Ort für jeden Schlüssel zu bestimmen und sein Vorhandensein in der Tabelle (oder der mit ihm verbundene Wert) kann durch Untersuchung dieser Zelle der Tabelle gefunden werden. Die offene Adressierung leidet jedoch unter Kollisionen, die auftreten, wenn mehr als ein Schlüssel derselben Zelle zugeordnet ist. Die Grundidee des Cuckoo-Hashings besteht darin, Kollisionen aufzulösen, indem zwei Hash-Funktionen anstelle von nur einer verwendet werden. Dadurch stehen für jeden Schlüssel zwei mögliche Positionen in der Hashtabelle zur Verfügung.

In einer der am häufigsten verwendeten Varianten des Algorithmus wird die Hash-Tabelle in zwei kleinere Tabellen gleicher Größe aufgeteilt und jede Hash-Funktion liefert einen Index in eine dieser beiden Tabellen. Es ist auch möglich, dass beide Hash-Funktionen Indizes in eine einzige Tabelle liefern.
