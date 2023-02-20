# Max Heap

Ein Heap (englisch wörtlich: Haufen oder Halde) in der Informatik ist eine zumeist auf Bäumen basierende abstrakte Datenstruktur. In einem Heap können Objekte oder Elemente abgelegt und aus diesem wieder entnommen werden. Sie dienen damit der Speicherung von Mengen. Den Elementen ist dabei ein Schlüssel zugeordnet, der die Priorität der Elemente festlegt. Häufig werden auch die Elemente selbst als Schlüssel verwendet.

Über die Menge der Schlüssel muss daher eine totale Ordnung festgelegt sein, über welche die Reihenfolge der eingefügten Elemente festgelegt wird. Beispielsweise könnte die Menge der ganzen Zahlen zusammen mit dem Vergleichsoperator < als Schlüsselmenge fungieren.

Der Begriff Heap wird häufig als bedeutungsgleich zu einem partiell geordneten Baum verstanden (siehe Abbildung). Gelegentlich versteht man einschränkend nur eine besondere Implementierungsform eines partiell geordneten Baums, nämlich die Einbettung in ein Feld (Array), als Heap.

Verwendung finden Heaps vor allem dort, wo es darauf ankommt, schnell ein Element mit höchster Priorität aus dem Heap zu entnehmen (HIFO-Prinzip), beispielsweise bei Vorrangwarteschlangen.

## Heap Bedingungen

Man unterscheidet Heaps in Min-Heaps und Max-Heaps. Bei Min-Heaps bezeichnet man die Eigenschaft, dass die Schlüssel der Kinder eines Knotens stets größer als oder gleich dem Schlüssel ihres Vaters sind, als Heap-Bedingung. Dies bewirkt, dass an der Wurzel des Baumes stets ein Element mit minimalem Schlüssel im Baum zu finden ist. Umgekehrt verlangt die Heap-Bedingung bei Max-Heaps, dass die Schlüssel der Kinder eines Knotens stets kleiner als oder gleich die ihres Vaters sind. Hier befindet sich an der Wurzel des Baumes immer ein Element mit maximalem Schlüssel.

Mathematisch besteht der Unterschied zwischen beiden Varianten nur in ihrer gegensätzlichen Ordnung der Elemente. Da die Definition von auf- und absteigend rein willkürlich ist, ist es also eine Auslegungsfrage, ob es sich bei einer konkreten Implementierung um einen Min-Heap oder um einen Max-Heap handelt.
