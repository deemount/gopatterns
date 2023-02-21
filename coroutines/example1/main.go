package main

import "log"

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger(resume chan int) int {
	return <-resume
}

func main() {
	resume := integers()
	log.Printf("generated: %d", generateInteger(resume))
	log.Printf("generated: %d", generateInteger(resume))
	log.Printf("generated: %d", generateInteger(resume))
	log.Printf("generated: %d", generateInteger(resume))
	log.Printf("generated: %d", generateInteger(resume))
	log.Printf("generated: %d", generateInteger(resume))
	log.Printf("%#v", resume)
}
