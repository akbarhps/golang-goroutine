package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestTimeAfter(t *testing.T) {
	channel := time.After(time.Second)
	time := <-channel
	fmt.Println(time)
}

func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(time.Second, func() {
		defer group.Done()
		fmt.Println(time.Now())
	})

	fmt.Println(time.Now())
	group.Wait()
}

func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(time.Second * 5)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTimeTick(t *testing.T) {
	ticker := time.Tick(time.Second)

	for time := range ticker {
		fmt.Println(time)
	}
}
