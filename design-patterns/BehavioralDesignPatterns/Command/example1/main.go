package main

import "fmt"

/*
	Das Beispiel wird Anhand eines Fernsehers betrachtet.

	Ein Fernsehgerät kann entweder eingeschaltet werden:

	* ON-Taste auf der Fernbedienung;
	* ON-Taste auf dem Fernseher selbst.

	Ich kann mit der Implementierung des ON-Befehlsobjekts mit dem Fernseher als Empfänger beginnen.
	Wenn die Ausführungsmethode für diesen Befehl aufgerufen wird, ruft sie ihrerseits die Funktion TV.on auf.
	Der letzte Teil ist die Definition eines Invokers. Ich werde eigentlich zwei Aufrufer haben:
		* die Fernbedienung und
		* den Fernseher selbst.

	Beide werden das ON-Befehlsobjekt einbetten.

	Beachte, dass ich dieselbe Anforderung in mehrere Aufrufer verpackt habe.
	So kann ich auch mit anderen Befehlen verfahren.

	Der Vorteil der Erstellung eines separaten Befehlsobjektes, ist die Entkopplung der
	der UI-Logik von der zugrunde liegenden Businesslogik. Es besteht keine Notwendigkeit
	unterschiedliche Handler für die einzelnen Aufrufer zu entwickeln. Das Befehlsobjekt enthält
	alle Informationen, die es zur Ausführung benötigt.

	Daher kann es auch für eine verzögerte Ausführung verwendet werden.
*/

// Invoker

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Command Interface

type Command interface {
	execute()
}

// Concrete Command

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// Concrete Command

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Receiver Interface

type Device interface {
	on()
	off()
}

// Concrete Receiver

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// Client code

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
