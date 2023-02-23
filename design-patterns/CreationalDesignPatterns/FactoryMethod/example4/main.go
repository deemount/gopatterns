package main

import (
	"fmt"
	"log"
)

// Abstrakt

type AbstraktDokument interface {
	Belegnummer() int
	setzeBelegnummer(belegnummer int)
	erzeugeDokument() string
}

// Abstraktion der Hierarchie-Ebene

type DokumentenTyp int

const (
	RECHNUNG DokumentenTyp = (iota + 1)
	GUTSCHRIFT
	LIEFERSCHEIN
)

func (typ DokumentenTyp) String() string {
	switch typ {
	case RECHNUNG:
		return "Rechnung"
	case GUTSCHRIFT:
		return "Gutschrift"
	case LIEFERSCHEIN:
		return "Lieferschein"
	}
	return ""
}

// Abstraktes Produkt

type Dokument struct {
	belegnummer int
}

func (dokument Dokument) Belegnummer() int {
	return dokument.belegnummer
}

func (dokument *Dokument) setzeBelegnummer(belegnummer int) {
	if belegnummer > 0 {
		dokument.belegnummer = belegnummer
	}
}

func (dokument *Dokument) erzeugeDokument() string {
	return fmt.Sprintf("dokument erzeugt (belegnummer: %d)", dokument.belegnummer)
}

// Rechnung (Ableitung von Dokument; konkretes Produkt)
type Rechnung struct {
	Dokument
}

// newRechnung (Erzeuger)
func newRechnung(belegnummer int) AbstraktDokument {
	return &Rechnung{
		Dokument: Dokument{
			belegnummer: belegnummer,
		},
	}
}

// Gutschrift (Ableitung von Dokument; konkretes Produkt)
type Gutschrift struct {
	Dokument
}

// newGutschrift (Erzeuger)
func newGutschrift(belegnummer int) AbstraktDokument {
	return &Gutschrift{
		Dokument: Dokument{
			belegnummer: belegnummer,
		},
	}
}

// Lieferschein (Ableitung von Dokument)
type Lieferschein struct {
	Dokument
}

// newLieferschein (Erzeuger)
func newLieferschein(belegnummer int) AbstraktDokument {
	return &Lieferschein{
		Dokument: Dokument{
			belegnummer: belegnummer,
		},
	}
}

type AbstraktFakturierer interface {
	AbstraktDokument
}

type Fakturierer struct{}

func newFakturierer(typ string, belegnummer int) AbstraktFakturierer {

	switch typ {
	case "Rechnung":
		return newRechnung(belegnummer)
	case "Gutschrift":
		return newGutschrift(belegnummer)
	case "Lieferschein":
		return newLieferschein(belegnummer)
	}

	return nil

}

// Execute

func main() {

	rechnung := newRechnung(50)
	log.Printf("main: rechnung %#v", rechnung.Belegnummer())

	gutschrift := newGutschrift(100)
	log.Printf("main: gutschrift %#v", gutschrift.Belegnummer())

	lieferschein := newLieferschein(200)
	log.Printf("main: lieferschein %#v", lieferschein.Belegnummer())

	var fakturierer AbstraktDokument

	fakturierer = newFakturierer(DokumentenTyp(1).String(), 10)
	log.Printf("main: fakturierer Belegenummer %d", fakturierer.Belegnummer())
	log.Printf("main: %s", fakturierer.erzeugeDokument())
	fakturierer = newFakturierer(DokumentenTyp(2).String(), 20)
	log.Printf("main: fakturierer Belegenummer %d", fakturierer.Belegnummer())
	log.Printf("main: %s", fakturierer.erzeugeDokument())
	fakturierer = newFakturierer(DokumentenTyp(3).String(), 30)
	log.Printf("main: fakturierer Belegenummer %d", fakturierer.Belegnummer())
	log.Printf("main: %s", fakturierer.erzeugeDokument())

}
