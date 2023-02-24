# Datenstrukturen

## (Singly) Linked List

In der Informatik ist eine verknüpfte Liste eine lineare Sammlung von Datenelementen, deren lineare Reihenfolge nicht durch ihre physische Anordnung im Speicher gegeben ist. Stattdessen verweist jedes Element auf das nächste. Es handelt sich um eine Datenstruktur, die aus einer Gruppe von Knoten besteht, die zusammen eine Folge darstellen. In der einfachsten Form besteht jeder Knoten aus Daten und einem Verweis (mit anderen Worten: einem Link) auf den nächsten Knoten in der Folge.

Eine lineare Datenstruktur ist eine Struktur, deren Elemente eine Art von Sequenz bilden.

### Warum spielt die physische Platzierung im Speicher keine Rolle?

Bei Arrays ist der Speicherplatz, den das Array einnimmt, fest, d. h. wenn ich ein Array mit 5 Elementen habe, nimmt die Sprache nur 5 Speicheradressen im Speicher auf, eine Adresse nach der anderen. Da diese Adressen eine Sequenz bilden, weiß das Array, in welchem Speicherbereich seine Werte gespeichert werden. Die physische Platzierung dieser Werte im Speicher erzeugt eine Sequenz.

Bei verknüpften Listen verhält es sich etwas anders. In der Definition heißt es, dass "jedes Element auf das nächste verweist", wobei "Daten und ein Verweis (mit anderen Worten: ein Link) auf den nächsten Knoten" verwendet werden. Das bedeutet, dass jeder Knoten der verknüpften Liste zwei Dinge speichert: einen Wert und einen Verweis auf den nächsten Knoten in der Liste.
