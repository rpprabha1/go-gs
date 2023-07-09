package models

var (
	WORKERS     int
	THREADS     int
	FOLDER_PATH string
)

type Configs struct {
	Gs string
}

type WorkerPool struct {
	maxWorkers int
	queuedTask chan func()
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		queuedTask: make(chan func()),
	}
}

func (wp *WorkerPool) AddTask(task func()) {
	wp.queuedTask <- task
}

func (wp *WorkerPool) Run() {
	go wp.run()
}

func (wp *WorkerPool) Done() {

}

func (wp *WorkerPool) run() {
	guard := make(chan struct{}, wp.maxWorkers)
	for task := range wp.queuedTask {
		guard <- struct{}{}
		go func() {
			task()
			<-guard
		}()
	}

}
