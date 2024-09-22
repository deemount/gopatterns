package main

/**

* Beispiel ohne Dependency Injection (hardcoded)

package main

import "fmt"

## Engine struct
type Engine struct{}

## Car struct mit einer festen Abhängigkeit von Engine
type Car struct {
    engine Engine
}

func NewCar() Car {
    return Car{engine: Engine{}}
}

func main() {
    car := NewCar()
    fmt.Println(car)
}

**/

import "fmt"

// Engine struct
type Engine struct{}

// Car struct, das die Abhängigkeit von außen erhält
type Car struct {
	engine Engine
}

// Konstruktor, der Engine injiziert
func NewCar(engine Engine) Car {
	return Car{engine: engine}
}

func main() {
	engine := Engine{}
	car := NewCar(engine) // Abhängigkeit wird beim Erstellen übergeben
	fmt.Println(car)
}
