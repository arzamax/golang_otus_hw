package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errorsCount struct {
	sync.Mutex
	value int
}

func (e *errorsCount) increment() {
	e.Lock()
	defer e.Unlock()
	e.value++
}

func (e *errorsCount) get() int {
	e.Lock()
	defer e.Unlock()
	return e.value
}

func taskWorker(nextTask chan Task, signal chan interface{}, errCount *errorsCount, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task := <-nextTask:
			if err := task(); err != nil {
				errCount.increment()
			}
		case <-signal:
			return
		}
	}
}

func Run(tasks []Task, n int, m int) error {
	if n <= 0 {
		n = len(tasks)
	}

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	var wg sync.WaitGroup
	errCount := errorsCount{value: 0}
	nextTask := make(chan Task)
	signal := make(chan interface{})

	wg.Add(n)
	defer wg.Wait()
	defer close(signal)

	for i := 0; i < n; i++ {
		go taskWorker(nextTask, signal, &errCount, &wg)
	}

	for _, task := range tasks {
		if value := errCount.get(); value >= m {
			return ErrErrorsLimitExceeded
		}

		nextTask <- task
	}

	return nil
}
