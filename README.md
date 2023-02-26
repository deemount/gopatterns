# gopatterns

Vorlagen für intuitives Arbeiten und für die alltägliche Ideengewinnung in Go. Vor einem Gebrauch mit 1 zu 1 Übernahme wird gewarnt. Die vorhandenen Beispiele sind dennoch alle getestet.

**Hinweis:**
Alle Beschreibungen sind in deutscher Sprache. Empfohlen wird nach ausgiebiger Studie die englische Sprache in der Verwendung.

## Entwurfsmuster

### Anzahl

Alle aufgelisteten Typen der Entwurfsmuster beinhalten ein oder mehrere Beispiele, die in der Anwendung zwar gestestet sind, einige dennoch ohne *_test.go-Datei dahergekommen. Diese werden nach und nach hinzungefügt. Außerdem sind in wenigen Beispielen keine ausführbare main-Funktion implementiert. Diese werden ebenfalls nach und nach in die Beispiele eingefügt.

* Verhaltensmuster: 11
* Nebenläufigkeitsmuster: 12
* Erzeugungsmuster: 6
* Strukturmuster: 8

***insgesamt:*** 37

### Übersicht

* Gang Of Four
  * Verhaltensmuster (Behavioral)
    * Chain Of Responsibility
    * Command
    * Interpreter
    * Iterator
    * Mediator
    * Memento
    * Observer
    * State
    * Strategy
    * Template Method
    * Visitor
  * Nebenläufigkeitsmuster (Concurrency)
    * Active Object
    * Balking
    * Binding Properties
    * Compute Kernel
    * Double Checked Locking
    * Event Based Asynchronous
    * Guarded Suspension
    * Join
    * Lock
    * Message Design Pattern
    * Monitor Object
    * Semaphores
  * Erzeugungsmuster (Creational)
    * AbstractFactory
    * Builder
    * FactoryMethod
    * Multiton
    * Prototype
    * Singleton
  * Strukturmuster (Structural)
    * Adapter
    * Aggregate
    * Bridge
    * Composite
    * Decorator
    * Facade
    * Flyweight
    * Proxy

## Datenstrukturen

* Binary Tree
* Doubly Linked List
* Gap Buffer
* Hash Table
* Heap
* Linked List
* Max Heap
* Min Heap
* Piece Table
* Priority Queue
* Rope
* Succinct
* Trie

## Weitere

Hier befinden sich weitere, nützliche Vorlagen bzw. Informationen, die nicht implizit zu den Entwurfsmuster und  Datenstrukturen dazugezählt werden (können).

* Algorithmen
  * Auswahl
    * KTH Smallest Number
  * Cache
    * Least Frequency Used (LFU)
    * Least Recently Used (LRU)
    * Most Recently Used (MRU)
  * Geometrie
    * Euklidische Distanz (Point Distance)
    * Hyperplane
    * Hypersphere
  * Lineare Funktionen
    * Schnittstellenberechnung von Linien
  * Hash
    * Cuckoo Hashing
    * Double Hashing
    * Hopscotch Hashing
    * Lineares Sondieren (Linear Probing)
    * Perfect Hash Function
    * Quadratic Probing
  * Mathematik
    * Ganzzahl Division (Integer Division)
  * Stringabgleich
    * Knut-Morris-Pratt
    * Most Frequent
  * Suchalgorithmen
    * Binary Search
    * Bi Section Method (Halbierungsmethode)
    * Exponential Search
    * Fractional Cascading
    * Ahnentafel (Genealogical Number System)
    * Interpolation Search
    * Linear Search
    * Noisy Binary Search
  * Baumstrukturen
    * Binary Space Partitioning (Binäre Raumaufteilung)
    * KDTree
    * Strahler Number
    * Quadtree
* Coroutine
  * Yield/Resume (Simluation von yield/resulme erfolgt über Nebenläufigkeit)
  * Generator (Simluation von yield/resulme erfolgt über Nebenläufigkeit)
* Protocoll Buffer
  * Node.js, Go und gRPC
* Templates
  * Nested
