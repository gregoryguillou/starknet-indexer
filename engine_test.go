package indexer

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	concurrently "github.com/tejzpr/ordered-concurrently/v3"
)

type workerFunction struct {
	number int
}

// This is an example of a worker function that simulates a 5 second task.
func (w *workerFunction) Run(ctx context.Context) interface{} {
	C := time.NewTicker(time.Second * 5)
	defer C.Stop()
	select {
	case <-ctx.Done():
		fmt.Println("Context done")
		return 0
	case <-C.C:
		fmt.Println("Expected output")
		return w.number
	}
}

type TestWorker struct{}

func (w *TestWorker) Start(ctx context.Context, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	worker := make(chan concurrently.WorkFunction, 2)
	outputs := concurrently.Process(ctx, worker, &concurrently.Options{PoolSize: 2, OutChannelBuffer: 2})

	go func() {
		for i := 0; i < 10; i++ {
			worker <- &workerFunction{
				number: i,
			}
			fmt.Println("message sent: ", i)
		}
		close(worker)
		fmt.Println("all workers have been scheduled and the channel closed")
	}()

	for output := range outputs {
		v, ok := output.Value.(int)
		if !ok {
			continue
		}
		fmt.Println(v)
	}
	fmt.Println("Done All")
}

func TestConcurrency(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go (&TestWorker{}).Start(ctx, wg)
	time.Sleep(time.Second * 10)
	cancel()
	wg.Wait()
	fmt.Println("Job cancelled")
}
