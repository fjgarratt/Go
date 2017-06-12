package main

import (
	"fmt"
	"sync"
)

const (
	NUM_VALUES_TO_SEND = 5

	NUM_RECEIVERS = 4
)

var (
	Values [NUM_VALUES_TO_SEND]int

	wgReceivers sync.WaitGroup
)

func Broadcast(chIn chan int, chOuts []chan int) {
	wgReceivers.Add(NUM_RECEIVERS)

	for i := 0; i < NUM_RECEIVERS; i++ {
		id := i
		go func(ch chan int) {
			for v := range ch {
				fmt.Printf("R%d =>: %d, ", id, v)
			}
			wgReceivers.Done()
		}(chOuts[i])
	}

	for v := range chIn {
		for _, c := range chOuts {
			c <- v
		}
	}

	for _, c := range chOuts {
		close(c)
	}
}

func main() {
	for idx := range Values {
		Values[idx] = idx + 1
	}

	chIn := make(chan int)

	chOuts := make([]chan int, NUM_RECEIVERS)
	for i := 0; i < NUM_RECEIVERS; i++ {
		chOuts[i] = make(chan int)
	}

	go Broadcast(chIn, chOuts)

	for i := 0; i < NUM_VALUES_TO_SEND; i++ {
		chIn <- i
		fmt.Println()
	}
	close(chIn)

	wgReceivers.Wait()

	fmt.Println()
}
