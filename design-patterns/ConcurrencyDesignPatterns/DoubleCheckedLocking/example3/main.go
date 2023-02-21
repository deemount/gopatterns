package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"
)

/*

	Hier haben wir ein Programm, das 4kb-Blöcke aus einer gemeinsamen Quelldatei liest
	und etwas damit macht. Das sieht vielleicht nach einer guten Idee aus,
	ist aber nicht so toll, wenn man bedenkt, was man davon hat. Startzeiten sind normalerweise
	kein Problem für vernetzte Anwendungen (obwohl sie vielleicht ein Problem für kurzlebige Dienste sind).

	Verwende dieses Muster für eine träge Initialisierung, wenn der Gewinn durch die Verzögerung
	der Initialisierung die zusätzlichen Wartungskosten wert ist.

	In den meisten Fällen sollte die Initialisierung zu Beginn des Programms mehr als genug sein.
	Verwende dieses Muster für träge Initialisierungen sparsam.

*/

type virtualFile struct {
	closed bool
	mu     sync.Mutex
	fd     *os.File
	Path   string
}

func (file *virtualFile) Read(p []byte) (n int, err error) {
	if file.fd == nil {
		file.mu.Lock()
		defer file.mu.Unlock()
		if file.fd == nil {
			f, err := os.Open(file.Path)
			if err != nil {
				return 0, err
			}
			file.fd = f
		}
	}
	return file.fd.Read(p)
}

func (file *virtualFile) Close() error {
	if file.closed {
		return nil
	}
	file.mu.Lock()
	defer file.mu.Unlock()
	if file.closed {
		return nil
	}
	err := file.Close()
	file.closed = true
	return err
}

func process(r io.Reader, done func()) {
	defer done()
	// do some other work
	time.Sleep(1 * time.Second)
	buf := make([]byte, 4*1024) // 4kb
	for {
		n, err := r.Read(buf)
		if err != nil {
			log.Printf("process: error %s", err)
		}
		log.Printf("process: %d", n)
	}
}

func main() {

	f := &virtualFile{Path: "var/log/byte-stream.log"}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go process(f, wg.Done)
		log.Printf("main: waitGroup %d", i)
	}
	wg.Wait()
}
