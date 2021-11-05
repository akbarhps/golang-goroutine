package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup
	counter := 0

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func() {
			once.Do(func() {
				counter++
			})
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
