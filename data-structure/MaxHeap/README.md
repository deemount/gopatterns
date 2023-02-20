# Max Heap

Man unterscheidet Heaps in Min-Heaps und Max-Heaps. Bei Min-Heaps bezeichnet man die Eigenschaft, dass die Schlüssel der Kinder eines Knotens stets größer als oder gleich dem Schlüssel ihres Vaters sind, als Heap-Bedingung. Dies bewirkt, dass an der Wurzel des Baumes stets ein Element mit minimalem Schlüssel im Baum zu finden ist. Umgekehrt verlangt die Heap-Bedingung bei Max-Heaps, dass die Schlüssel der Kinder eines Knotens stets kleiner als oder gleich die ihres Vaters sind. Hier befindet sich an der Wurzel des Baumes immer ein Element mit maximalem Schlüssel.

Mathematisch besteht der Unterschied zwischen beiden Varianten nur in ihrer gegensätzlichen Ordnung der Elemente. Da die Definition von auf- und absteigend rein willkürlich ist, ist es also eine Auslegungsfrage, ob es sich bei einer konkreten Implementierung um einen Min-Heap oder um einen Max-Heap handelt.
