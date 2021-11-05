package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
	}()

	data := <-channel
	fmt.Println(data)
	// fmt.Println(<-channel) this work to

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyOut(channel)
	go OnlyIn(channel)

	time.Sleep(3 * time.Second)
}
