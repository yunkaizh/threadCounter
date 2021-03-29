package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var cnt int
var stop int32

func work() {
	// Do nothing..
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}

func count() {
	var loop = true
	for loop == true {
		// Check whether to terminate the loop
		if atomic.LoadInt32(&stop) == 1 {
			break
		}

		// update the num of go routines
		var cur int = runtime.NumGoroutine()
		if cur > cnt {
			cnt = cur
		}
	}
}

func main() {
	var numGo int = 1000000
	var err error

	// Get the number of GO routinues to create
	if len(os.Args) >= 2 {
		numGo, err = strconv.Atoi(os.Args[1])
		if err != nil {
			numGo = 1000000
		}
	}
	fmt.Println("Creating ", numGo, " routines")

	cnt = 0
	// Kick off count() to count the max go routines
	go count()

	for i := 0; i < numGo; i++ {
		wg.Add(1)
		go work()
	}

	// Wait until all go routines to finish
	wg.Wait()

	// Change the stop indicator to 1 to terminal count()
	atomic.SwapInt32(&stop, 1)
	fmt.Println("Max go routines - ", cnt)
	// Use 0 to return the previous setting.
	fmt.Println("Max number of CPUs go can utilize - ", runtime.GOMAXPROCS(0))
	fmt.Println("Num of logical CPU - ", runtime.NumCPU())
}
