package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/sachinnagesh/go-worker-pool/entity"
)

func main() {
	log.Info("### Go Worker Pool ###")

	// create tasks
	tasks := make([]entity.Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = entity.Task{ID: i + 1}
	}

	// create worker pool
	wp := entity.WorkerPool{
		Tasks:       tasks,
		Concurrency: 5, // number of workers that can run at a time
	}

	// run the pool
	wp.Run()
	log.Info("All tasks have been processed!!!")

}
