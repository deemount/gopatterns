# Dependency Injection

Dependency Injection ist ein Verhaltensmuster. Es regelt, wie Objekte Abhängigkeiten erhalten und wie diese verwaltet werden, anstatt sie selbst zu erzeugen oder zu steuern.

Dependency Injection (DI) gehört zur Kategorie der Inversions of Control (IoC).

## Zusammenhang mit Inversion of Control (IoC)

DI ist ein spezieller Mechanismus, um das Prinzip der Inversion of Control umzusetzen. IoC beschreibt allgemein die Verlagerung der Kontrolle über das Erstellen und Verwalten von Objekten weg von der Klasse selbst hin zu einem externen Mechanismus oder Framework. Bei DI erfolgt dies, indem die Abhängigkeiten (Dependencies) von außen in die Klassen eingespeist werden.

## DI als Verhaltensmuster

Dependency Injection wird als Verhaltensmuster klassifiziert, weil es den Ablauf und die Art der Kommunikation zwischen Objekten beeinflusst. Anstatt dass eine Klasse selbst entscheidet, welche Abhängigkeiten sie benötigt und wie diese erstellt werden, werden diese Abhängigkeiten von außen bereitgestellt. Dies verändert das Verhalten der Klasse in Bezug auf die Abhängigkeitsverwaltung und -nutzung.

In Go (Golang) gibt es kein spezielles Framework für Dependency Injection (DI) wie in Sprachen wie Java oder C#, da Go eine andere Philosophie hinsichtlich Einfachheit und expliziten Abhängigkeiten verfolgt. Dennoch kann das Konzept von Dependency Injection in Go auf eine manuelle und weniger formalisierte Weise angewendet werden.

## Dependency Injection in Go

In Go wird Dependency Injection in der Regel durch Konstruktor-Injection oder Parameter-Injection erreicht, bei der Abhängigkeiten als Argumente an Funktionen oder Strukturen übergeben werden. Im Gegensatz zu Frameworks, die DI automatisch handhaben (wie Spring in Java), geschieht dies in Go explizit und manuell.

## Unterschiede zwischen Go und typischen DI-Frameworks (wie in Java)

1. Kein automatisiertes DI-Framework: In Go gibt es keine automatisierten Mechanismen oder Frameworks wie Spring in Java, die den Lebenszyklus von Objekten und deren Abhängigkeiten managen. Alles wird explizit gehandhabt, was der Philosophie von Go entspricht (explizite statt implizite Magie).

2. Explizite Abhängigkeitsverwaltung: In Go erfolgt die Abhängigkeitsverwaltung oft durch Konstruktoren oder Funktionen. Entwickler müssen alle Abhängigkeiten manuell injizieren, was zu einer klareren Struktur führen kann, aber auch etwas mehr Boilerplate-Code erfordert.

3. Keine Reflexion oder Annotationen: Go verwendet keine Reflexion oder Annotationen, um Abhängigkeiten zu erkennen und zu injizieren. In Java oder C# können DI-Frameworks wie Spring oder .NET auf diese Mechanismen zurückgreifen, um die Abhängigkeitsverwaltung zu automatisieren.

4. Testbarkeit und Einfachheit: Ähnlich wie in anderen Sprachen erleichtert auch in Go die manuelle DI den Unit-Testing-Prozess, da Abhängigkeiten leicht durch Mock-Objekte ersetzt werden können. Allerdings geschieht das in Go durch direkte Konstruktion der Abhängigkeiten.

## Beispiel für andere IoC-Umsetzungsformen

Service Locator: Ein weiteres IoC-Muster, bei dem eine Klasse ihre Abhängigkeiten durch eine zentrale Registrierungsstelle (Service Locator) abruft. Im Gegensatz zur DI, wo die Abhängigkeiten direkt injiziert werden, fragt die Klasse hier aktiv nach ihren Abhängigkeiten.

## Fazit

Dependency Injection ist ein Verhaltensmuster und eine spezielle Form der Inversion of Control. Es ermöglicht eine flexible und testbare Strukturierung von Code, indem es die Verwaltung der Abhängigkeiten von der Klasse selbst entkoppelt und diese Verantwortung nach außen verlagert.

Dependency Injection in Go unterscheidet sich stark von Framework-gesteuerten Ansätzen, die man in anderen Sprachen findet. In Go ist die DI sehr explizit und einfach gehalten, ohne zusätzlichen Overhead durch Frameworks. Das macht den Code oft übersichtlicher, erfordert aber mehr manuelle Arbeit beim Verwalten der Abhängigkeiten.
