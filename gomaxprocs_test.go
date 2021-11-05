package golanggoroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("totalCpu:", totalCpu)

	runtime.GOMAXPROCS(8)
	maxProcs := runtime.GOMAXPROCS(-1)
	fmt.Println("maxProcs:", maxProcs)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("totalGoroutine:", totalGoroutine)
}
