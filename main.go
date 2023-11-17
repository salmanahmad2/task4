package main

import (
	"fmt"
	"sync"
)

var sharedBuffer []byte
var mutex sync.Mutex

type Routine struct {
	M int
	N int
}

func main() {
	sharedBuffer = make([]byte, 20)

	routines := []Routine{
		{M: 8, N: 2},
		{M: 8, N: 8},
		{M: 8, N: 16},
		{M: 2, N: 8},
	}

	for _, r := range routines {
		runRoutine(r)
	}
}

func runRoutine(r Routine) {
	for i := 0; i < r.M; i++ {
		go reader(i)
	}

	for j := 0; j < r.N; j++ {
		go writer(j)
	}

	select {}
}

func writer(wIndex int) {
	for {
		mutex.Lock()
		index := wIndex % len(sharedBuffer)
		sharedBuffer[index] = byte(wIndex)
		fmt.Println("written to buffer: ", wIndex)
		mutex.Unlock()
	}
}

func reader(rIndex int) {
	for {
		mutex.Lock()
		index := rIndex % len(sharedBuffer)
		value := sharedBuffer[index]
		fmt.Println("read from buffer: ", value)
		mutex.Unlock()
	}
}
