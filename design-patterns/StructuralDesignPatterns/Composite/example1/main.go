package main

import "fmt"

/*
	Lassen Sie mich versuchen, das Composite-Muster anhand eines Beispiels aus dem Dateisystem eines Betriebssystems zu verstehen.
	Im Dateisystem gibt es zwei Arten von Objekten: Dateien und Ordner. Es gibt Fälle, in denen Dateien und
	Ordner auf die gleiche Weise behandelt werden sollten. Hier kommt das Composite-Muster ins Spiel.

	Ich stelle mir vor, dass ich eine Suche nach einem bestimmten Schlüsselwort in meinem Dateisystem durchführen muss.
	Dieser Suchvorgang gilt sowohl für Dateien als auch für Verzeichnisse.
	Bei einer Datei wird nur der Inhalt der Datei durchsucht; bei einem Ordner werden alle Dateien des Ordners
	nach dem Schlüsselwort durchsucht.
*/

// Leaf

type Component interface {
	search(string)
}

// Component interface

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

// Composite

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

// Client code

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
