# Double Checked Locking

Doppelt geprüftes Sperren lässt träge Initialisierungen idempotent und thread-safe machen. Das klingt zwar auf dem Papier gut, aber ist es in der Praxis wirklich so praktisch?

Wofür wird die träge Initialisierung verwendet:

* Aufschieben der Initialisierung von teuren Ressourcen, bis sie tatsächlich benötigt werden (Beispiel 1)
* Aufschieben der Initialisierung, bis bestimmte Fakten feststehen (Beispiel 3)

Das Double-Check-Locking stellt sicher, dass nur ein Minimum an Arbeit anfällt und dass keine doppelten Initialisierungen stattfinden.

## Hinweis

Das sync-Paket, hat einen Typ namens sync.Once, der das Problem der doppelten Initialisierung beheben soll und nach Möglichkeit anstelle von Double Checked Locking verwendet werden sollte.

Wie Donald Knuth sagte:

***"Vorzeitige Optimierung ist die Wurzel allen Übels"***

Double Checked Locking ist ein ziemlich fortschrittliches Entwurfsmuster. Und obwohl es auf dem Papier wirklich gut klingt, würde ich davon abraten, es sei denn, wenn Tausende oder sogar Millionen solcher Objekte konstruiert und gebündelt werden und die Kosten für die Konstruktion ein Flaschenhals sind.

Doppelt geprüftes Sperren schadet der Lesbarkeit Ihres Codes. Verwende dieses Muster erst, wenn es tatsächlich der Sicherheit oder Leistung dient.

Es ist jedoch erwähnenswert, dass das Double-Checked-Locking-Entwurfsmuster in Go allgemein als Anti-Muster betrachtet wird, da es schwierig sein kann, es richtig zu machen und zu subtilen Gleichzeitigkeitsfehlern führen kann. In den meisten Fällen ist es besser, einfachere Muster wie Lazy Initialization mit sync.Once zu verwenden oder Dependency Injection einzusetzen, um sicherzustellen, dass eine einzige Instanz einer Struktur in allen Teilen der Anwendung gemeinsam genutzt wird.

## Beispiel

### Träge Initialisierung (Einfach)

```go
func (obj *object) method() {
  if obj.property == nil {
    obj.property = newProperty()
  }
}
```

### Träge Initialisierung (Thread safe)

```go
func (obj *object) method() {
  if obj.property == nil {
    obj.mu.Lock()
    obj.property = newProperty()
    obj.mu.Unlock()
  }
}
```

### Träge Initialisierung (Double checked; Thread safe)

Ich nehme ein hypothetisches Beispiel für einen Fall, in dem zwei Goroutinen diese Methode gleichzeitig aufrufen. Beide erreichen das Prädikat check obj.property == nil und versuchen, die Sperre zu erhalten. An diesem Punkt gewinnt eine von ihnen und initialisiert die Eigenschaft, während die andere Goroutine wartet. Sobald die Mutex entsperrt ist, kommt die zweite Goroutine und initialisiert die Eigenschaft erneut. Durch das Hinzufügen einer zweiten Prüfung stelle ich sicher, dass die nachfolgende Goroutine dieselbe Eigenschaft nicht erneut initialisiert

```go
func (obj *object) method() {
  if obj.property == nil {
    obj.mu.Lock()
      if obj.property == nil {
        obj.property = newProperty()
      }
      obj.mu.Unlock()
  }
}
```
