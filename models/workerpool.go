package models

import "sync"

type WorkerPool struct {
	maxWorkers int
	queuedTask chan func()
	wg         sync.WaitGroup
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		queuedTask: make(chan func()),
	}
}

func (wp *WorkerPool) AddTask(task func()) {
	wp.wg.Add(1)
	wp.queuedTask <- task
}

func (wp *WorkerPool) Run() {
	go wp.run()
}

func (wp *WorkerPool) Done() {
	wp.wg.Wait()
}

func (wp *WorkerPool) run() {
	guard := make(chan struct{}, wp.maxWorkers)
	for task := range wp.queuedTask {
		guard <- struct{}{}
		go func() {
			task()
			<-guard
			wp.wg.Done()
		}()
	}
}
