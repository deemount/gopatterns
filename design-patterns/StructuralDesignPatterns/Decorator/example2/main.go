package main

import "log"

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)
		result := fn(n)
		log.Println("Execution is completed with the result", result)
		return result
	}
}

func Double(n int) int {
	return n * 2
}

func main() {

	f := LogDecorate(Double)
	f(5)

}
