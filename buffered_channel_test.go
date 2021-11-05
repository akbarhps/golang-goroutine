package golanggoroutine

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Hello"
	channel <- "World"
	channel <- "!"

	fmt.Println(<-channel) // Hello
	fmt.Println(<-channel) // World
	fmt.Println(<-channel) // !

	fmt.Println(cap(channel)) // 3
	fmt.Println(len(channel)) // 0
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Perulangan ke %d", i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for counter := 0; counter < 2; {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2: ", data)
			counter++
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for counter := 0; counter < 2; {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2: ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
	}
}
