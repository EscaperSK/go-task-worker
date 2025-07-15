package worker

import (
	"log"
	"math/rand"
	"tasks/database"
)

type Task struct {
	Id          int
	Duration    int
	Probability int
}

func (t *Task) start() {
	database.UpdateTask(t.Id, database.Processing)
}

func (t *Task) finishOrFail() {
	taskIsProcessed := rand.Intn(100) < t.Probability

	if taskIsProcessed {
		database.UpdateTask(t.Id, database.Processed)
		log.Printf("Task #%d finished.\n", t.Id)
	} else {
		database.UpdateTask(t.Id, database.New)
		log.Printf("Task #%d failed.\n", t.Id)
	}
}
