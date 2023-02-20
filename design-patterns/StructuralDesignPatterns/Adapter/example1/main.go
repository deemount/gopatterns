package main

import "fmt"

/*
	Ich habe einen Client-Code, der einige Funktionen eines Objekts (Lightning-Port) erwartet,
	aber ich habe ein anderes Objekt namens Adaptee (Windows-Laptop), das die
	die gleiche Funktionalität bietet, aber über eine andere Schnittstelle (USB-Anschluss)

	An dieser Stelle kommt das Adapter-Muster ins Spiel. Ich erstelle eine Struktur
	Typ, bekannt als Adapter, der:

	* sich an die gleiche Schnittstelle hält, die der Client erwartet (Lightning-Anschluss).
	* Die Anfrage vom Client an den Adapter in der Form übersetzt, die der Adapter erwartet.

	Der Adapter akzeptiert einen Lightning-Stecker und übersetzt dann dessen Signale in ein
	ein USB-Format und leitet sie an den USB-Anschluss des Windows-Laptops weiter.
*/

// Client code

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// Client interface

type Computer interface {
	InsertIntoLightningPort()
}

// Service

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Unknown Service

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// Adapter

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

// Client code

func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
