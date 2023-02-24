package main

import "fmt"

/*

 */

func perfectHash(data []string) func(string) int {
	size := len(data)
	hash := make(map[string]int, size)

	// First hash function: returns the ASCII value of the first character
	for i := 0; i < size; i++ {
		hash[string(data[i][0])] = i
	}

	// Second hash function: returns the length of the string
	secondHash := func(str string) int {
		return len(str)
	}

	// Third hash function: returns the sum of the ASCII values of the characters
	thirdHash := func(str string) int {
		sum := 0
		for _, char := range str {
			sum += int(char)
		}
		return sum
	}

	// Generate a candidate hash function until a perfect hash function is found
	var candidate func(string) int
	for {
		candidate = func(str string) int {
			return (hash[string(str[0])] + secondHash(str) + thirdHash(str)) % size
		}
		used := make(map[int]bool, size)
		collision := false
		for _, str := range data {
			index := candidate(str)
			if used[index] {
				collision = true
				break
			}
			used[index] = true
		}
		if !collision {
			break
		}
	}

	return candidate
}

func main() {
	data := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "orange", "peach", "quince", "raspberry", "strawberry", "tangerine", "watermelon"}
	hash := perfectHash(data)
	for _, str := range data {
		fmt.Printf("%s: %d\n", str, hash(str))
	}
}
