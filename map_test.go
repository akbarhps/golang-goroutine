package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func StoreToMap(group *sync.WaitGroup, data *sync.Map, value int) {
	group.Add(1)
	defer group.Done()

	data.Store(value, value)
}

func TestStoreToMap(t *testing.T) {
	data := sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go StoreToMap(&group, &data, i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
}
