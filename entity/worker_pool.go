package entity

import "sync"

// worker pool definition
type WorkerPool struct {
	Tasks       []Task
	Concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// funciton to execute worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	//initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	//start workers
	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	//Send tasks to the task channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	// wait for all tasks to finish

	wp.wg.Wait()

}
