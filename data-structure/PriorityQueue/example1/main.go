package main

import (
	"container/heap"
	"fmt"
)

/*

	Beispiel:

	In diesem Beispiel definiere ich eine Task-Struktur, die eine Aufgabe mit einem
	Namen und einer Priorität darstellt und einen TaskQueue-Typ, der ein Slice von
	Task-Zeigern ist. Anschließend werden die sort.Interface-Methoden für TaskQueue
	definiert, die zur Implementierung des Verhaltens der Prioritätswarteschlange verwendet
	werden. Ich definiere zudem Methoden für TaskQueue, die es mir ermöglichen, Tasks zu
	schieben und zu entfernen sowie die Priorität einer Task zu aktualisieren.

	In der Hauptfunktion initialisiere ich dann eine neue TaskQueue, füge der Warteschlange
	einige Aufgaben hinzu, aktualisiere die Priorität einer der Aufgaben und entferne dann
	die Aufgaben in der Reihenfolge ihrer Priorität aus der Warteschlange.

*/

// Define a task struct
type Task struct {
	name     string
	priority int
	index    int
}

// Define a task queue (priority queue)
type TaskQueue []*Task

// Implement the sort.Interface for TaskQueue
func (tq TaskQueue) Len() int {
	return len(tq)
}
func (tq TaskQueue) Less(i, j int) bool {
	return tq[i].priority > tq[j].priority
}
func (tq TaskQueue) Swap(i, j int) {
	tq[i], tq[j] = tq[j], tq[i]
	tq[i].index = i
	tq[j].index = j
}

// Define methods for TaskQueue
func (tq *TaskQueue) Push(x interface{}) {
	n := len(*tq)
	task := x.(*Task)
	task.index = n
	*tq = append(*tq, task)
}
func (tq *TaskQueue) Pop() interface{} {
	old := *tq
	n := len(old)
	task := old[n-1]
	task.index = -1 // for safety
	*tq = old[0 : n-1]
	return task
}

// Define a function to update the priority of a task
func (tq *TaskQueue) update(task *Task, priority int) {
	task.priority = priority
	heap.Fix(tq, task.index)
}

func main() {
	// Initialize a new task queue
	tasks := make(TaskQueue, 0)

	// Add some tasks to the queue
	heap.Push(&tasks, &Task{"task1", 5, 0})
	heap.Push(&tasks, &Task{"task2", 1, 0})
	heap.Push(&tasks, &Task{"task3", 3, 0})

	// Update the priority of task1
	tasks.update(tasks[0], 2)

	// Pop tasks off the queue in priority order
	for tasks.Len() > 0 {
		task := heap.Pop(&tasks).(*Task)
		fmt.Printf("name: %s, priority: %d\n", task.name, task.priority)
	}
}
