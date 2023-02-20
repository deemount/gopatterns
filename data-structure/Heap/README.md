# Heap

Ein Heap (englisch wörtlich: Haufen oder Halde) in der Informatik ist eine zumeist auf Bäumen basierende abstrakte Datenstruktur. In einem Heap können Objekte oder Elemente abgelegt und aus diesem wieder entnommen werden. Sie dienen damit der Speicherung von Mengen. Den Elementen ist dabei ein Schlüssel zugeordnet, der die Priorität der Elemente festlegt. Häufig werden auch die Elemente selbst als Schlüssel verwendet.

Über die Menge der Schlüssel muss daher eine totale Ordnung festgelegt sein, über welche die Reihenfolge der eingefügten Elemente festgelegt wird. Beispielsweise könnte die Menge der ganzen Zahlen zusammen mit dem Vergleichsoperator < als Schlüsselmenge fungieren.

Der Begriff Heap wird häufig als bedeutungsgleich zu einem partiell geordneten Baum verstanden. Gelegentlich versteht man einschränkend nur eine besondere Implementierungsform eines partiell geordneten Baums, nämlich die Einbettung in ein Feld (Array), als Heap.

## Anwendungsbeispiele

Verwendung finden Heaps vor allem dort, wo es darauf ankommt, schnell ein Element mit höchster Priorität aus dem Heap zu entnehmen (HIFO-Prinzip), beispielsweise bei Vorrangwarteschlangen.

Einige der wichtigsten Eigenschaften, die Heaps nützlich machen, sind:

* Wir können den maximalen oder minimalen Wert einer gegebenen Menge von "n" Werten in O(1) Zeit finden. Bei Verwendung einer auf einem Array basierenden Liste würde dies normalerweise O(n) Zeit in Anspruch nehmen.
* Das Hinzufügen und Entfernen von Werten nimmt O(log(n)) Zeit in Anspruch, wobei die Max- oder Min-Heap-Eigenschaft erhalten bleibt. Dies wird beim Sortieren einer Liste von Elementen mittels Heap-Sortierung verwendet, bei der jedes Element einzeln in einen Heap eingefügt wird.