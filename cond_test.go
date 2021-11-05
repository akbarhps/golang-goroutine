package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait() // wait for signal

	fmt.Println("Done", value)

	cond.L.Unlock()
}

func TestWaitCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			cond.Signal() // send signal to wait condition
		}

		// time.Sleep(time.Second)
		// cond.Broadcast() // send signal to all wait condition
	}()

	group.Wait()
}
