package workerpool

import (
	"strconv"
	"testing"
	"time"
)

func TestWorkerPool_ProcessTasks(t *testing.T) {
	wp := NewWorkerPool()
	defer wp.Shutdown()

	wp.AddWorker()
	wp.AddWorker()
	wp.AddWorker()
	wp.AddWorker()
	wp.AddWorker()

	for i := 1; i <= 50; i += 1 {
		task := "task" + strconv.Itoa(i)
		wp.AddTask(task)
	}
	time.Sleep(1 * time.Millisecond)
}
