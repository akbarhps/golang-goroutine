package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	group.Add(1)
	defer group.Done()

	fmt.Println("Hello")
	time.Sleep(time.Second)
}

func TestWaitGroup(t *testing.T) {
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go RunAsynchronous(&group)
	}

	group.Wait()
	fmt.Println("Complete")
}
