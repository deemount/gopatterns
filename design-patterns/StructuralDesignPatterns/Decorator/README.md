# Decorator

Decorator ist ein strukturelles Entwurfsmuster, das es mir ermöglicht, bestehenden Objekten dynamisch neues Verhalten hinzuzufügen, indem Ich sie in spezielle Wrapper-Objekte einbette. Diese Wrapper fungieren als Erweiterung meines Codes, die zur Wiederverwendung von Funktionen beitragen und den bestehenden Code nicht verändert.

## Beziehung zu anderen Mustern

* Composite und Decorator haben ähnliche Strukturdiagramme, da sie beide auf rekursiver Komposition basieren, um eine unbestimmte Anzahl von Objekten zu organisieren.

* Ein Decorator ist wie ein Composite, hat aber nur eine untergeordnete Komponente. Es gibt einen weiteren wichtigen Unterschied: Decorator fügt dem umhüllten Objekt zusätzliche Aufgaben hinzu, während Composite lediglich die Ergebnisse seiner Kinder "zusammenfasst". Muster können jedoch auch hilfreich sein: Ich kann den Decorator verwenden, um das Verhalten eines bestimmten Objekts im Composite-Baum zu erweitern.

* Mit dem Decorator kann ich die Struktur eines Objekts ändern, während Ich mit der Strategie sein Innenleben veränderen kann.

* Decorator und Proxy haben ähnliche Strukturen, aber sehr unterschiedliche Zwecke. Beide Muster beruhen auf dem Prinzip der Komposition, wonach ein Objekt einen Teil der Arbeit an ein anderes delegieren sollte. Der Unterschied besteht darin, dass ein Proxy normalerweise den Lebenszyklus seines Dienstobjekts selbst verwaltet, während die Komposition des Decorators immer vom Client kontrolliert wird.

### Pro

* Das Verhalten eines Objekts/Strukts kann erweitert/geändert werden, ohne eine neue Struktur zu erstellen.
* Hinzufügen oder Entfernen von Verantwortlichkeiten eines Objekts zur Laufzeit.
* Ermöglicht die Kombination mehrerer Verhaltensweisen durch Umhüllen eines Objekts mit mehreren Dekoratoren.

### Contra

* Es ist schwierig, einen bestimmten Wrapper aus dem Wrapper-Stack zu entfernen.
* Es ist schwierig, einen Dekorator so zu implementieren, dass sein Verhalten nicht von der Reihenfolge im Dekoratorstapel abhängt.
* Der anfängliche Konfigurationscode für die Schichten kann hässlich aussehen.
