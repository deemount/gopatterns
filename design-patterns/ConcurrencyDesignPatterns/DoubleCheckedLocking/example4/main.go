package main

import (
	"sync"
)

/*

	Beispiel:

	In dieser Implementierung wird die Variable once vom Typ sync.Once verwendet,
	um sicherzustellen, dass die Initialisierung des Singleton-Objekts nur einmal
	durchgeführt wird. Die Funktion GetInstance() ruft die Methode Do() von once auf
	und übergibt eine Funktion, die die Instanzvariable nur dann initialisiert, wenn sie
	noch nicht initialisiert wurde. Dadurch wird sichergestellt, dass das Singleton-Objekt
	nur einmalig erstellt wird, d. h. nur dann, wenn es zum ersten Mal benötigt wird.

	Die Instanzvariable wird als Variable auf Paketebene deklariert, so dass sie von jedem
	Teil des Pakets aus zugänglich ist. Sie ist anfangs auf null gesetzt und wird auf das
	neu erstellte Singleton-Objekt gesetzt, wenn GetInstance() zum ersten Mal aufgerufen wird.

	Diese Implementierung stellt sicher, dass das Singleton-Objekt nur einmal erstellt wird,
	selbst in einer nebenläufigen Umgebung, da die Do()-Methode von once garantiert nur einmal
	ausgeführt wird. Dadurch wird die Race Condition vermieden, die auftreten könnte, wenn
	mehrere Goroutinen gleichzeitig versuchen würden, das Singleton-Objekt zu erstellen.

*/

type Singleton struct {
	// Fields of the singleton object
}

var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		// This function will only be executed once,
		// no matter how many times GetInstance() is called.
		instance = &Singleton{
			// Initialize fields of the singleton object
		}
	})
	return instance
}
