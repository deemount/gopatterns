package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

/*
	Wichtig:
	Es muss eine Funktion definiert werden, die verwendet wird, um den
	den Worker-Prozess zu definieren. Diese Funktion erhält einen Kanal,
	um die Daten zu verarbeiten.

	Außerdem müssen die Goroutinen für die Worker gestartet werden, bevor
	die Daten an den Kanal übergeben werden.
	Werden die Daten an den Kanal übergeben, beginnen die Worker mit der
	Verarbeitung der Werte.

	Die Übergabe der Werte erfolgt in den Zeilen 64 bis 67 sehen können.
*/

type city struct {
	name     string
	location string
}

func createCity(record city) {
	time.Sleep(1 * time.Millisecond)
}

func readData(cityChn chan []city) {
	var cities []city

	csvFile, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		cities = append(cities, city{name: line[0], location: line[1]})
	}
	cityChn <- cities

	close(cityChn)

}

func worker(cityChn chan city) {
	for val := range cityChn {
		createCity(val)
	}
}

func main() {

	startTime := time.Now()
	cities := make(chan []city)

	go readData(cities)

	// more workers decreases the total time
	const workers = 10
	jobs := make(chan city, 1000)

	for w := 1; w <= workers; w++ {
		go worker(jobs)
	}

	counter := 0
	for _, val := range <-cities {
		counter++
		jobs <- val
	}

	fmt.Println("records saved:", counter)
	fmt.Println("total time:", time.Since(startTime))

}
