package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const PROCS = 4                 // Constant for number of processes to break into
const NUM_ELEM = 1000000        // Number of elements in slice
const BLOCKS = NUM_ELEM / PROCS // Size of each block determined by PROCS and NUM_ELEM

var glob_sum float64 = 0
var m sync.Mutex

func make_average(w *sync.WaitGroup, aSlice []float64) {

	var summer float64 = 0         // Define the container for the result
	for _, value := range aSlice { // loop over items in the slice
		summer += value // add them to the running total
	}

	w.Done()
	m.Lock()                                    // signal to the waitgroup this process is done
	glob_sum += (summer / float64(len(aSlice))) // add the result into the channel
	m.Unlock()
}

func main() {

	var wg sync.WaitGroup // Define WaitGroup to check all processes have summed before continuing
	//var ch = make(chan float64) // Make a float64 channel to receive each process' result
	wg.Add(PROCS) // Define the number of processes in the waitgroup

	mil := make([]float64, NUM_ELEM) // Make a slice for the number of data points

	for i := 0; i < NUM_ELEM; i++ { // Populate the data points with a random number between 1 and 1,000,000
		mil[i] = rand.Float64() * 1000000
	}

	start_point := 0             // Define the starting point for the first block of data
	for i := 0; i < PROCS; i++ { // Step over the data in BLOCKS size lumps
		go make_average(&wg, mil[start_point:(start_point+BLOCKS)]) // Launch goprocess, passing waitgroup, slice and channel for result
		start_point += BLOCKS                                       // Increment the starting point to the start of the next block
	}

	wg.Wait() // Wait until all the processes have completed

	//var external float64 = 0 // Define the outside sum for the grand total

	/*
		for a := 0; a < PROCS; a++ {
				external += <-ch // Extract results from the channel
		}*/

	//close(ch)

	fmt.Println(glob_sum / PROCS) // Display the result divided by number of processes
}
