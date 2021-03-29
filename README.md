# The threadCounter
The threadCounter.go is implemented to count the max go routines as well as to count the max CPUs Go Scheduler uses.

## Number of OS Threads
Based on my understanding, GO scheduler uses one M (OS Thread) per P (Processor) to schedule G (Goroutines). 
Unless there is a blocking syscall, Go scheduler does not spawn new OS Thread. In the event of a blocking syscall, 
the blocked OS Thread is detached and a new OS Thread is brought in to handle other Go routines. The max number of
OS Thread to run user-level code will not exceed GOMAXPROC. GOMAXPROC is used to estimate the max OS Threads used by Go Scheduler.

## Number of Go routines
The runtime.NumGoroutine is used to estimate the max number of live go routines.


## Usage
By default, it will create 100K routines. <br/>
`go run threadCounter.go [number of go routines to create]`<br/>
E.g.,<br/>
`go run threadCounter.go`<br/>
`go run threadCounter.go 1000000`<br/>
