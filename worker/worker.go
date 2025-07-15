package worker

import (
	"log"
	"time"
)

var queue chan struct{}

func Run(size int) {
	queue = make(chan struct{}, size)
}

func AddTask(task *Task) {
	if queue == nil {
		log.Fatal("Worker is not initialized.")
	}

	task.start()

	go func() {
		queue <- struct{}{}
		go func() {
			processTask(task)
			<-queue
		}()
	}()
}

func processTask(task *Task) {
	duration := max(time.Duration(task.Duration), 0)

	time.Sleep(time.Millisecond * duration)

	task.finishOrFail()
}
