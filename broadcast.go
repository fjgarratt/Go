package main

import (
	"fmt"
	"time"
)

var channnelsBroadcast [3]chan string

func ChatRoom(idx int, chBroadcasted chan string) {
	for msg := range chBroadcasted {
		fmt.Printf("%d got: %s\n", idx, msg)
	}
}

func Broadcaster(chMessages chan string) {
	for msg := range chMessages {
		for _, chChat := range channnelsBroadcast {
			chChat <- msg
		}
	}
}

func main() {
	//create an array of channels
	for idx, _ := range channnelsBroadcast {
		channnelsBroadcast[idx] = make(chan string)
	}

	// make a main channel
	chanMain := make(chan string)

	// range over the channels launching a goprocess for each channel passing the index and the channel
	for idx, _ := range channnelsBroadcast {
		go ChatRoom(idx, channnelsBroadcast[idx])
	}

	// launch the main broadcast channel
	go Broadcaster(chanMain)

	// define the time to stop (in 5 seconds)
	stopAtTime := time.Now().Add(5 * time.Second)

	for {
		// put a message into the main channel
		chanMain <- time.Now().String() + " bientÃ´t la pause !"
		// sleep for 1 second
		time.Sleep(1 * time.Second)

		// check to see if we should stop
		if time.Now().After(stopAtTime) {
			close(chanMain)
			// close the channels
			for _, ch := range channnelsBroadcast {
				close(ch)
			}
			break // break out of the for loop because we've finished
		}
	}

	fmt.Println("Done")
}
