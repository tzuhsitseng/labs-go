package concurrency

import (
	"fmt"
	"sync"
)

type Worker interface {
	Start()
	Dispatch(job Job)
	Shutdown()
}

type MathWorker struct {
	jobChan   []chan Job
	closeChan chan struct{}
}

const (
	defaultWorkers = 10
	defaultBuffers = 10
)

func NewMathWorker() *MathWorker {
	jobChan := make([]chan Job, defaultWorkers)
	for i := range jobChan {
		jobChan[i] = make(chan Job, defaultBuffers)
	}

	return &MathWorker{
		jobChan:   jobChan,
		closeChan: make(chan struct{}),
	}
}

func (w *MathWorker) Start() {
	for i := range w.jobChan {
		go func(idx int, c <-chan Job) {
			for job := range c {
				executor, err := NewExecutor(job)
				if err != nil {
					fmt.Println("new executor failed", err)
				}

				if err = executor.DoJob(job); err != nil {
					fmt.Println("execute job failed", err)
				}

				fmt.Printf("executed job by worker[%d]\n", idx)
			}

			fmt.Printf("close worker[%d]\n", idx)
			w.closeChan <- struct{}{}
		}(i, w.jobChan[i])
	}
}

func (w *MathWorker) Dispatch(jobs ...Job) {
	for _, job := range jobs {
		id := job.GetID()
		w.jobChan[id%len(w.jobChan)] <- job
	}
}

func (w *MathWorker) Shutdown() {
	wg := sync.WaitGroup{}
	wg.Add(len(w.jobChan))

	fmt.Println("start to shutdown all workers")

	go func() {
		for {
			select {
			case <-w.closeChan:
				wg.Done()
			}
		}
	}()

	for i := range w.jobChan {
		close(w.jobChan[i])
	}

	wg.Wait()
	fmt.Println("all workers are already closed")
}
