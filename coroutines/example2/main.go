package main

import "fmt"

func Count(start int, end int) chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// Blocks on the operation
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}

func main() {
	fmt.Println("Kein Bier im Schrank")

	for i := range Count(1, 99) {
		fmt.Println("Nehme Flasche aus der Kiste.", i, "Flaschen im Schrank")
	}

	fmt.Println(100, "Flaschen sind im Schrank")
}
