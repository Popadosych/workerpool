package workerpool

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type WorkerPool struct {
	taskCh       chan string
	nextID       atomic.Int64
	stopCh       chan struct{}
	workersCount sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		taskCh: make(chan string),
		stopCh: make(chan struct{}),
	}
}

func (wp *WorkerPool) startWorker(id int64) {
	wp.workersCount.Add(1)
	defer wp.workersCount.Done()
	go func() {
		for {
			select {
			case task := <-wp.taskCh:
				fmt.Printf("Worker %d processing task: %s\n", id, task)
			case <-wp.stopCh:
				return
			}
		}
	}()
}

func (wp *WorkerPool) AddWorker() {
	wp.startWorker(wp.nextID.Add(1))
}

func (wp *WorkerPool) RemoveWorker() {
	wp.stopCh <- struct{}{}
}

func (wp *WorkerPool) AddTask(task string) {
	wp.taskCh <- task
}

func (wp *WorkerPool) Shutdown() {
	close(wp.stopCh)
	wp.workersCount.Wait()
}
