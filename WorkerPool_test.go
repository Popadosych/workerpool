package workerpool

import (
	"strconv"
	"testing"
	"time"
)

func TestWorkerPool_ProcessTasks(t *testing.T) {
	wp := NewWorkerPool()

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

	wp.Shutdown()

	if wp.nextID.Load() != 51 {
		t.Errorf("Not all tasks done")
	}
}
