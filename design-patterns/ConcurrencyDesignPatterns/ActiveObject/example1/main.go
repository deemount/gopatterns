package main

import (
	"fmt"
	"log"
)

/*
	In diesem Beispiel definiere ich eine Schnittstelle für das Active Object mit zwei Methoden:
   	Do() und Stop().

   	Die Methode Do() nimmt eine Aktion und einen Wert als Argumente entgegen und führt die
   	angegebene Methode in einer Goroutine aus. Die Methode Stop() hält das Active Object an.

   	Ich definiere dann die activeObject-Struktur mit zwei Kanälen:
   		* actions zum Ausführen der Methode und
   		* quit zum Beenden des Active Object.

   	In der Funktion NewActiveObject() erstelle ich ein neues ActiveObject und starte es als Goroutine.

   	Die Do()-Methode erstellt eine Schließung zur Ausführung der angegebenen Methode basierend auf den
	Argumenten action und value. Anschließend sendet sie die Closure an den actions-Kanal,
	um vom aktiven Objekt ausgeführt zu werden.

   	Schließlich sendet dann die Methode Stop() eine Nachricht an den Quit-Kanal, um das Aktive Objekt zu beenden.
*/

// Define an interface for the Active Object
type ActiveObject interface {
	Do(action string, value interface{}) // Method to be executed by the Active Object
	Stop()                               // Method to stop the Active Object
}

// Define the Active Object structure
type activeObject struct {
	actions chan func()   // Channel for executing the method
	quit    chan struct{} // Channel to stop the Active Object
}

// Create a new Active Object
func NewActiveObject() ActiveObject {

	obj := &activeObject{
		actions: make(chan func()),
		quit:    make(chan struct{}),
	}

	// Start the Active Object as a goroutine
	go func() {
		for {
			select {
			case action := <-obj.actions:
				action()
			case <-obj.quit:
				return
			}
		}
	}()

	return obj
}

// Implement the Do() method of the Active Object
func (a *activeObject) Do(action string, value interface{}) {
	// Create a closure to execute the method
	fn := func() {
		switch action {
		case "Print":
			msg := value.(string)
			fmt.Println(msg)
		case "Add":
			nums := value.([]int)
			sum := 0
			for _, num := range nums {
				sum += num
			}
			fmt.Println(sum)
		}
	}

	// Send the closure to the actions channel
	a.actions <- fn
}

// Implement the Stop() method of the Active Object
func (a *activeObject) Stop() {
	// Send a message to the quit channel to stop the Active Object
	a.quit <- struct{}{}
}

func main() {

	obj := NewActiveObject()
	obj.Do("Add", []int{100})
	obj.Do("Print", "this is a message")
	obj.Do("Add", []int{10})
	obj.Stop()

	log.Printf("main: %#v", obj)

}
