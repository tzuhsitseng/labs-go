package main

import (
	"sync"
	"time"

	"github.com/tzuhsitseng/labs-go/concurrency"
)

func main() {
	worker := concurrency.NewMathWorker()
	worker.Start()
	worker.Dispatch(
		concurrency.AddJob{Left: 1, Right: 2},
		concurrency.AddJob{Left: 3, Right: 4},
		concurrency.AddJob{Left: 5, Right: 6},
		concurrency.AddJob{Left: 7, Right: 8},
		concurrency.AddJob{Left: 9, Right: 10},
		concurrency.SubtractJob{Left: 10, Right: 5},
		concurrency.SubtractJob{Left: 8, Right: 2},
		concurrency.SubtractJob{Left: 7, Right: 1},
		concurrency.SubtractJob{Left: 4, Right: 3},
		concurrency.SubtractJob{Left: 9, Right: 6},
	)

	wg := sync.WaitGroup{}
	wg.Add(1)
	ticker := time.After(time.Second * 3)
	go func() {
		select {
		case <-ticker:
			worker.Shutdown()
			wg.Done()
		}
	}()
	wg.Wait()
}
