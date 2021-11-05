package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{ // override the default sync.Pool
		New: func() interface{} {
			return "Busy"
		},
	}

	pool.Put("Hello")
	pool.Put("World")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
