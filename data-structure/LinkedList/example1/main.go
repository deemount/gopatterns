package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	Beispiel:

	Für ein soziales Netzwerk lasse ich mich von Twitter inspirieren und
	erstelle einen Post-Typ, der einen Body, ein PublishDate und einen Link
	zum nächsten Post hat.

	Die Feed-Struktur hat einen Anfang (oder Start), der auf den ersten Beitrag
	im Feed verweist und eine Längeneigenschaft, die die Größe des Feeds zu
	einem bestimmten Zeitpunkt speichert.

	Dann wird eine Append-Funktion für den Feed-Typ erstellt. Die Funktion erzeugt
	einen Zeiger auf einen Post-Wert, wobei das Argument body als Body des Posts dient
	und das publishDate auf die Unix-Timestamp-Darstellung der aktuellen Zeit gesetzt wird.
	Dann prüft die Funktion, ob die Länge des Feeds 0 ist - das bedeutet, dass der Feed
	keine Beiträge hat und der erste hinzugefügte Beitrag als Startbeitrag festgelegt
	werden sollte, der praktischerweise Start heißt.

	Wenn der Feed leer ist, wird der neue Beitrag als Anfang und Ende des Feeds festgelegt,
	andernfalls wird der neue Beitrag als Endelement festgelegt und an den vorherigen Beitrag
	in der Liste angehängt. Zusätzlich zu seiner Einfachheit hat dieser Algorithmus nun
	eine Big-O-Komplexität von O(1), auch bekannt als "konstante Zeit". Das bedeutet,
	dass Append unabhängig von der Länge der Feed-Struktur die gleiche Leistung erbringt.

	Verknüpfte Listen können aufgrund ihres Aufbaus, bei dem jedes Element einen Zeiger auf den
	nächsten Knoten in der Liste hat, über den gesamten Speicherplatz verstreut sein. Das bedeutet,
	dass das Hinzufügen/Entfernen von Knoten aus der Liste aus Sicht des Speicherplatzes günstig ist.
	Wenn man einen Knoten aus einer verknüpften Liste entfernen möchte, müssen nur die Nachbarn des
	entfernten Knotens verknüpft werden, und das ist es.

	Garbage-Collecting-Sprachen (wie Go) machen dies sogar noch einfacher, da wir uns nicht um die
	Freigabe des zugewiesenen Speichers kümmern müssen - der GC wird eingreifen und alle unbenutzten
	Objekte entfernen.

	Die Funktion Remove nimmt das Veröffentlichungsdatum eines Beitrags als Argument,
	anhand dessen sie erkennt, welcher Beitrag gelöscht (oder entlinkt) werden muss.
	Die Funktion ist recht klein. Wenn sie feststellt, dass das Startelement des Feeds entfernt werden
	soll, ordnet sie einfach den Start des Feeds dem zweiten Beitrag im Feed neu zu.
	Andernfalls springt sie durch alle Beiträge im Feed, bis sie auf einen Beitrag stößt,
	dessen Veröffentlichungsdatum mit dem als Funktionsargument übergebenen übereinstimmt.
	Wird ein solcher gefunden, werden einfach der vorherige und der nächste Beitrag im Feed miteinander
	verbunden, wobei der mittlere (übereinstimmende) Beitrag aus dem Feed entfernt wird.

	Die Funktion Insert ist ein sehr ähnlicher Algorithmus wie der in der Funktion Remove, denn obwohl beide
	einen ganz anderen Zweck erfüllen (Hinzufügen bzw. Entfernen eines Beitrags im Feed), basieren beide auf
	einem Suchalgorithmus.

	Das bedeutet, dass beide Funktionen den gesamten Feed durchlaufen und nach einem Beitrag suchen,
	dessen Veröffentlichungsdatum mit dem im Argument der Funktion angegebenen übereinstimmt.
	Der einzige Unterschied besteht darin, dass Insert den neuen Beitrag tatsächlich an der Stelle einfügt,
	an der die Daten übereinstimmen, während Remove den Beitrag aus dem Feed entfernt.

	Die Hauptfunktion erzeugt einen Zeiger auf eine Feed-Struktur, zwei Post-Strukturen
	mit einigen Dummy-Inhalten. Diese ruft zweimal die Append-Funktion im Feed auf,
	was dazu führt, dass der Feed eine Länge von 2 hat.

	Ich untersuche die beiden Werte des Feeds, indem ich auf den Anfang des Feeds
	und dem nächsten Element, das der zweite Post ist, zugreife.

*/

type Post struct {
	body        string
	publishDate int64 // Unix timestamp
	next        *Post
}

type Feed struct {
	length int
	start  *Post
	end    *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length++
}

func (f *Feed) Remove(publishDate int64) {
	if f.length == 0 {
		panic(errors.New("Feed is empty"))
	}

	var previousPost *Post
	currentPost := f.start

	for currentPost.publishDate != publishDate {
		if currentPost.next == nil {
			panic(errors.New("No such Post found."))
		}

		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next

	f.length--
}

func (f *Feed) Insert(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		var previousPost *Post
		currentPost := f.start

		for currentPost.publishDate < newPost.publishDate {
			previousPost = currentPost
			currentPost = previousPost.next
		}

		previousPost.next = newPost
		newPost.next = currentPost
	}
	f.length++
}

func (f *Feed) Inspect() {
	if f.length == 0 {
		fmt.Println("Feed is empty")
	}
	fmt.Println("========================")
	fmt.Println("Feed Length: ", f.length)

	currentIndex := 0
	currentPost := f.start

	for currentIndex < f.length {
		fmt.Printf("Item: %v - %v\n", currentIndex, currentPost)
		currentPost = currentPost.next
		currentIndex++
	}
	fmt.Println("========================")
}

func main() {
	rightNow := time.Now().Unix()
	f := &Feed{}
	p1 := &Post{
		body:        "Lorem ipsum",
		publishDate: rightNow,
	}
	p2 := &Post{
		body:        "Dolor sit amet",
		publishDate: rightNow + 10,
	}
	p3 := &Post{
		body:        "consectetur adipiscing elit",
		publishDate: rightNow + 20,
	}
	p4 := &Post{
		body:        "sed do eiusmod tempor incididunt",
		publishDate: rightNow + 30,
	}

	f.Append(p1)
	f.Append(p2)
	f.Append(p3)
	f.Append(p4)

	f.Inspect()

	newPost := &Post{
		body:        "This is a new post",
		publishDate: rightNow + 15,
	}
	f.Insert(newPost)
	f.Inspect()
}
