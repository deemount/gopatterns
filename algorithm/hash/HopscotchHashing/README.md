# Hopscotch Hashing

## Konzept

Hopscotch-Hashing ist ein Schema in der Computerprogrammierung zur Auflösung von Hash-Kollisionen von Werten von Hash-Funktionen in einer Tabelle mit offener Adressierung. Es eignet sich auch gut für die Implementierung einer gleichzeitigen Hashtabelle. Das Hopscotch-Hashing wurde 2008 von Maurice Herlihy, Nir Shavit und Moran Tzafrir eingeführt. Der Name leitet sich von der Abfolge der Hops ab, die den Einfügealgorithmus der Tabelle charakterisieren.

Der Algorithmus verwendet ein einzelnes Array mit n Buckets. Für jeden Bereich ist seine Nachbarschaft eine kleine Sammlung von H aufeinanderfolgenden Bereichen (d. h. Bereichen mit Indizes, die nahe am ursprünglichen Hash-Bereich liegen). Die gewünschte Eigenschaft der Nachbarschaft ist, dass die Kosten für das Auffinden eines Elements in den Buckets der Nachbarschaft nahe an den Kosten für das Auffinden im Bucket selbst liegen (z. B. dadurch, dass die Buckets in der Nachbarschaft in dieselbe Cache-Zeile fallen). Die Größe der Nachbarschaft muss so bemessen sein, dass sie im schlimmsten Fall eine logarithmische Anzahl von Elementen aufnehmen kann (d. h. sie muss log(n) Elemente aufnehmen können), im Durchschnitt jedoch nur eine konstante Anzahl. Wenn die Nachbarschaft eines Buckets voll ist, wird die Größe der Tabelle angepasst.

Beim Hopscotch-Hashing wird, wie beim Cuckoo-Hashing und anders als bei der linearen Suche, ein bestimmtes Element immer in der Nachbarschaft seines Hash-Buckets eingefügt und gefunden. Mit anderen Worten, es wird immer entweder in seinem ursprünglichen Hash-Element oder in einem der nächsten H-1 benachbarten Elemente gefunden. H könnte z.B. 32 sein, eine übliche Maschinenwortgröße. Die Nachbarschaft ist also ein "virtueller" Bereich, der eine feste Größe hat und sich mit den folgenden H-1 Bereichen überschneidet.

Um die Suche zu beschleunigen, enthält jeder Bucket (Array-Eintrag) ein "Hop-Informations"-Wort, eine H-Bit-Bitmap, die angibt, welche der nächsten H-1-Einträge Elemente enthalten, die zum virtuellen Bucket des aktuellen Eintrags gehasht wurden. Auf diese Weise kann ein Element schnell gefunden werden, indem man sich das Wort ansieht, um festzustellen, welche Einträge zu dem Bereich gehören, und dann die konstante Anzahl von Einträgen durchsucht (die meisten modernen Prozessoren unterstützen spezielle Bitmanipulationsoperationen, die das Nachschlagen in der "hop-information"-Bitmap sehr schnell machen).