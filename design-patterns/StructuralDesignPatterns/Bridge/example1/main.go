package main

import "fmt"

/*
	Ich habe zwei Arten von Computer: Mac und Windows.
	Außerdem zwei Arten von Drucker: Epson und HP.

	Beide Computer und Drucker müssen in jeder Kombination miteinander funktionieren.
	Der Kunde möchte sich nicht um die Details des Anschlusses von Druckern
	an die Computer kümmern.

	Wenn ich neue Drucker einführe, möchte ich nicht, dass unser Code exponentiell wächst.
	Anstatt vier Strukturen für die 2*2-Kombination zu erstellen, erstelle ich
	zwei Hierarchien:

	* Abstraktionshierarchie für die Computer
	* Implementierungshierarchie für die Drucker

	Diese beiden Hierarchien kommunizieren miteinander über eine Brücke, wobei die
	Abstraktion (Computer) einen Verweis auf die Implementierung (Drucker) enthält.
	Sowohl die Abstraktion als auch die Implementierung können unabhängig voneinander
	entwickelt werden ohne sich gegenseitig zu beeinträchtigen.
*/

// Abstraction

type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Refined abstraction

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Implementation

type Printer interface {
	PrintFile()
}

// Concrete implementation

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

// Client code

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
