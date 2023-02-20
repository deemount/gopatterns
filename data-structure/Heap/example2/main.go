package main

import (
	"container/heap"
	"log"
	"strings"
	"time"
)

// Each holiday has a name and a date
type Holiday struct {
	name string
	date time.Time
}

// for convenience, we can create a string representation to print the holiday instance
func (h Holiday) String() string {
	return "[" + h.name + ", " + h.date.Format("Jan 2") + "]"
}

// We can create a new type to represent a list of holidays
type Holidays []Holiday

// Implement the container/heap interface
func (h Holidays) Len() int { return len(h) }
func (h Holidays) Less(i, j int) bool {
	return strings.Compare(h[i].name, h[j].name) < 0
}
func (h Holidays) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Holidays) Push(x interface{}) {
	*h = append(*h, x.(Holiday))
}

func (h *Holidays) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	holidays := &Holidays{}
	heap.Init(holidays)

	heap.Push(holidays, Holiday{name: "Christmas", date: time.Date(2023, time.December, 25, 0, 0, 0, 0, time.Local)})
	heap.Push(holidays, Holiday{name: "Labour Day", date: time.Date(2023, time.May, 1, 0, 0, 0, 0, time.Local)})
	heap.Push(holidays, Holiday{name: "Diwali", date: time.Date(2023, time.November, 23, 0, 0, 0, 0, time.Local)})

	log.Println("heap: holidays (before swap) ", holidays)

	holidays.Swap(1, 0)

	log.Println("heap: holidays (after swap) ", holidays)

}
