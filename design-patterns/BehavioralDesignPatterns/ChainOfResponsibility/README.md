# Chain Of Responsibility

Chain of Responsibility ist ein Verhaltensmuster, das die Weitergabe von Anfragen entlang der Kette potenzieller Handler ermöglicht, bis einer von ihnen die Anfrage bearbeitet.

Das Muster erlaubt es mehreren Objekten, die Anfrage zu bearbeiten, ohne die Senderklasse an die konkreten Klassen der Empfänger zu koppeln. Die Kette kann dynamisch zur Laufzeit mit jedem Handler zusammengestellt werden, der einer Standard-Handler-Schnittstelle folgt.
