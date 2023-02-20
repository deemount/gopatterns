package main

import "fmt"

/*
	Hinweis:
	Das Entwurfsmuster wird am Beispiel einer Krankenhausanwendung erläutert

	Ein Krankenhaus könnte mehrere Abteilungen haben, wie z.B.:

	* Rezeption
	* Arzt
	* Medizinischer Untersuchungsraum
	* Kassierer

	Wenn ein Patient ankommt, geht er zuerst

	* zur Rezeption,
	* dann zum Arzt,
	* dann zum Medizinischen Untersuchungsraum,
	* und dann zum Kassierer

	Der Patient wird durch eine Kette von Abteilungen geschickt, wobei jede Abteilung den Patienten
	jede Abteilung den Patienten in der Kette weiterschickt, sobald ihre Funktion abgeschlossen ist.

	Das Muster ist anwendbar, wenn es mehrere Kandidaten gibt, die dieselbe Anfrage bearbeiten.
	Es ist auch nützlich, wenn ich nicht möchte, dass der Client den Empfänger auswählt, da es
	mehrere Objekte die Anfrage bearbeiten können. Ein weiterer nützlicher Fall ist, wenn ich den Client von den Empfängern entkoppeln möchte
	der Client von den Empfängern entkoppelt werden soll - der Client muss dann nur das erste Element in der Kette kennen.

	Wie im Beispiel des Krankenhauses geht ein Patient zunächst zur Rezeption. Dann, basierend auf dem
	Status des Patienten sendet die Rezeption an den nächsten Handler in der Kette.
*/

// Handler Interface

type Department interface {
	execute(*Patient)
	setNext(Department)
}

// Concrete

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

// Client code

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
