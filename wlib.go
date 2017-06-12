package main

import (
	"fmt"
	"sync"
	"time"

	"./mylib"
)

func incrementVals(c mylib.Connections_t, vlu string, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	c[vlu] = c[vlu] + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println(mylib.Connections)

	var calls = [2]string{"usa", "eur"}

	for _, vlu := range calls {
		go incrementVals(mylib.Connections, vlu, &wg)
	}

	wg.Wait()

	fmt.Println(mylib.Connections)
}
