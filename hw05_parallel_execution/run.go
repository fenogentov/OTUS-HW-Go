package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 1 {
		n = 1
	}
	if m < 1 {
		m = len(tasks) + 1
	}

	var wg sync.WaitGroup
	//	chQuit := make(chan struct{})
	inJob := make(chan Task)

	errCounter := int32(m) //nolint

	// горутина распределения заданий
	wg.Add(1)
	go func(tasks []Task) {
		defer wg.Done()
		for _, job := range tasks {
			if atomic.LoadInt32(&errCounter) <= 0 {
				break
			}
			inJob <- job
		}
		close(inJob)
	}(tasks)

	// горутины выполнения заданий
	for j := 0; j < n; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				j, ok := <-inJob
				if !ok {
					return
				}
				e := j()
				if e != nil {
					atomic.AddInt32(&errCounter, -1)
				}
			}
		}()
	}
	wg.Wait()

	if errCounter <= 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
