# Active Object

Das Active Object Design Pattern ist ein Concurrency Pattern, das die Methodenausführung vom Methodenaufruf entkoppelt, um die Nebenläufigkeit und Skalierbarkeit zu verbessern. In Go kann dieses Entwurfsmuster mit Kanälen und Goroutinen implementiert werden.

Eine nützliche Verfeinerung der einschränkenden Goroutine ist etwas, das unter vielen Namen bekannt ist, aber der wahrscheinlich häufigste: Das ***Active Object*** Concurrency Pattern. Es wird häufig in Serversoftware mit vielen gleichzeitigen Verbindungen verwendet, bei denen der Overhead der Verwendung eines Threads für jede Verbindung zu groß ist.

Das ***Active Object*** in Go kann durch eine Goroutine implementiert werden, die Closures aus einem Channel (chan func()) liest und ausführt. Dieses einfache System bedeutet, dass alle Closures, die den Code enthalten, der auf die Daten zugreift, in derselben Goroutine in der Reihenfolge ausgeführt werden, in der sie in den Kanal gestellt werden.

## Problembeschreibung

In Go ist es sehr einfach, versehentlich von verschiedenen Goroutinen aus auf dieselben Daten zuzugreifen, wodurch eine Race Condition entsteht. Konventionell vermieden dieser Art von Problem wird mit einem sog. Mutex. In Go wird das mit sync.Mutex bewerkstelligt.

Eine oft bessere und in Go bevorzugte Möglichkeit besteht darin, den gleichzeitigen Zugriff auf die Daten durch verschiedene Goroutinen zu vermeiden. Eine Möglichkeit besteht darin, Nachrichten (über einen Kanal) an eine eingeschränkte Goroutine zu senden, die für den gesamten Zugriff und die Kontrolle der Daten verantwortlich ist. Durch die Beschränkung des gesamten Zugriffs auf eine einzige Goroutine sind keine Sperren erforderlich.

Ich kann die Verwendung eines Wertes auch einschränken, indem ich ihn jeweils nur innerhalb einer Goroutine verwende. Dies geschieht in Go idiomatisch, indem die Kontrolle über die Variable durch einen Kanal übertragen wird. Außerdem gibt es "Low-Level"-Methoden, um "sperrfreie" Parallelität zu erreichen, indem man atomare Operationen verwendet.

## Sperren (Deadlocks) vermeiden

Was ist also der Vorteil der Vermeidung von Sperren (sog. Deadlocks)? Die Verwendung von Mutexen in komplexen Szenarien ist berüchtigt für die Aufrechterhaltung subtiler Race Conditions oder potenzieller Deadlocks oder einfach nur für Leistungsprobleme, wie das unbeabsichtigte Halten einer Sperre beim Blockieren von Eingabe und Ausgabe.

Der Streit um Sperren wird oft als großer Vorteil der Vermeidung von Sperren angeführt, aber das ist nicht wirklich das Problem. Schließlich werden bei der Verwendung einer einschränkenden Go-Routine (wie oben beschrieben) die Sperren durch die Verwendung der Go-Routine ersetzt.

Tatsächlich ist die richtige Verwendung von Sperren oft effizienter als die Verwendung von Kanälen; es ist nur so, dass dies in der Regel komplizierten und/oder sich wiederholenden Code erfordert.

Ein typisches Szenario für die Verwendung von Schlössern (sog. locks) ist zum Beispiel folgendes:

* lock
* check "something"
* unlock
* do a lengthy operation based on knowledge obtained at 2
* lock
* check "something" again (repeating test at 2)
* update data using results at 4 (if "something" hasn't changed)
* unlock

Es gibt jedoch einige Leistungsvorteile

Erstens: Wenn viele Threads auf Mutex(s) blockieren, wird der Overhead für das Thread-Switching wichtig. Auch wenn die Zeit für das Thread-Switching nur in Mikrosekunden gemessen wird, kann sich das bei Tausenden von Threads summieren. Natürlich wird dieser Effekt durch die Verwendung von Go-Routinen gemildert, da Go sie auf eine kleine Anzahl von Threads verteilt, aber er kann immer noch erheblich sein.

Außerdem denke ich, dass Go versucht, Go-Routinen jedes Mal auf demselben Kern laufen zu lassen, was bedeutet, dass eine begrenzte Go-Routine besser sein kann, um den CPU-Cache zu erhalten, was große Leistungsvorteile haben könnte.

## Stichpunkte

* Entkopplung der Methodenausführung vom Methodenaufruf
* Vereinfacht den synchronisierten Zugriff auf Objekte, die sich in ihren eigenen "thread of control" befinden
* Wird häufig in verteilten Systemen verwendet, die Multi-Thread-Server erfordern
* Wird oft mit Message Passing implementiert
* Ähnlich wie das Actor-Muster. Im Actor-Muster ist jedes Objekt ein aktives Objekt
