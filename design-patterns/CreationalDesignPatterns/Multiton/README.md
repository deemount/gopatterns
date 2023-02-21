# Multiton

## Konzept

Das Multiton-Entwurfsmuster ist ein Erstellungsmuster, das sicherstellt, dass eine Klasse nur eine begrenzte Anzahl von Instanzen hat, wobei jede Instanz durch einen von mehreren möglichen Schlüsseln dargestellt wird. Der Schlüssel wird als Bezeichner verwendet, um eine einzelne Instanz aus einer Liste von Instanzen zu erhalten, die vom Multiton verwaltet werden.

Das Multiton-Entwurfsmuster ähnelt dem Singleton-Muster, allerdings mit einem gewissen Unterschied. Das Singleton-Muster stellt sicher, dass eine Klasse nur eine Instanz hat. Im Gegensatz dazu sorgt das Multiton-Muster dafür, dass eine Klasse eine begrenzte Anzahl von Instanzen hat, die jeweils einen eigenen eindeutigen Bezeichner haben. Das Multiton-Muster ist hilfreich in Situationen, in denen mehrere Instanzen einer Klasse erforderlich sind, aber jede Instanz eine eindeutige Identität haben muss.

## Zielsetzung

* Bereitstellung einer einzigen Instanz einer Klasse für jeden Schlüssel, wobei der Zugriff auf jede Instanz über den Schlüssel erfolgt.
* Erzwingen eines einzigen Kontrollpunkts für die Verwaltung von Instanzen einer Klasse.
* Unterstützung eines effizienten Instanzabrufs unter Verwendung eines Schlüssels.
* Um sicherzustellen, dass nur eine begrenzte Anzahl von Instanzen einer Klasse erstellt wird.
* Um die Notwendigkeit zu vermeiden, mehrere Instanzen einer Klasse in verschiedenen Teilen des Codes zu erstellen.

### Pro

* Erzwingt eine einzige Instanz einer Klasse pro Schlüssel
* Bietet einen zentralen Kontrollpunkt für die Instanzen
* Unterstützt effizientes Auffinden von Instanzen über den Schlüssel
* Es vermeidet die Notwendigkeit, mehrere Instanzen einer Klasse in verschiedenen Teilen des Codes zu erstellen

### Contra

* Führt zu einer erhöhten Kopplung zwischen dem Client-Code und der Multiton-Registry
* Macht es schwierig, die Implementierung der Multiton-Registry zu ändern, da der Client-Code von ihrem spezifischen Verhalten abhängen kann
* Es kann zu unnötigem Speicherverbrauch führen, wenn Instanzen für Schlüssel erstellt werden, die nie verwendet werden
* Es ist schwierig, die Anzahl der Instanzen einer bestimmten Klasse zu bestimmen
