package main

import "log"

/*

	In diesem Beispiel definiere ich eine Worker-Struktur, die Eigenschaften wie taskType, Priorität und maxTasks enthält.
	Außerdem definiere ich eine Task-Struktur, die einen TaskType, eine Priorität und ein Datenfeld enthält.

	Die Funktion workerPool erstellt einen Kanal zur Aufnahme eingehender Aufgaben,
	startet eine feste Anzahl von Worker-Goroutinen und fügt dem Kanal einige Beispielaufgaben hinzu.
	Jeder Worker prüft die Eigenschaften der eingehenden Aufgabe und verarbeitet sie, wenn sie mit seinen eigenen
	Eigenschaften übereinstimmt.

	Durch die Verwendung der Bindungseigenschaften auf diese Weise, kann ich einen flexibleren und effizienteren Worker-Pool erstellen,
	der verschiedene Arten von Aufgaben mit unterschiedlichen Prioritäten und Verarbeitungsanforderungen verarbeiten kann.

*/

type Worker struct {
	id       int
	taskType string
	priority int
	maxTasks int
}

func (w *Worker) process(task string) {
	// Process the task
	log.Printf("worker: processing %s", task)
}

func NewWorker(id int, taskType string, priority int, maxTasks int) *Worker {
	return &Worker{id, taskType, priority, maxTasks}
}

type Task struct {
	taskType string
	priority int
	data     interface{}
}

func NewTask(taskType string, priority int, data interface{}) *Task {
	log.Println("task: init new")
	return &Task{taskType, priority, data}
}

func workerPool(numWorkers int, workerProps []*Worker) {
	// Create a channel to hold incoming tasks
	tasks := make(chan *Task)

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		log.Printf("workerpool: start worker goroutine %d", i)
		go func(id int) {
			// Find the worker with the matching properties
			var worker *Worker
			for _, w := range workerProps {
				if w.id == id {
					worker = w
					break
				}
			}

			// Process tasks until the channel is closed
			for task := range tasks {
				log.Printf("workerpool: range task value %v", task)
				if task.taskType == worker.taskType && task.priority >= worker.priority {
					// Process the task if it matches the worker's properties
					worker.process(task.data.(string))
				}
			}
		}(i)
	}

	log.Println("workerpool: eof range goroutines")

	// Add tasks to the channel
	tasks <- NewTask("type1", 1, "task1")
	tasks <- NewTask("type2", 2, "task2")

	log.Printf("workerpool: show %#v", &tasks)

	// Close the channel when all tasks have been added
	close(tasks)
}

func main() {

	workerPool(2, []*Worker{
		{id: 1, taskType: "type1", priority: 1, maxTasks: 1},
		{id: 1, taskType: "type1", priority: 1, maxTasks: 1},
		{id: 2, taskType: "type2", priority: 3, maxTasks: 1},
	})

}
