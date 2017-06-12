package main

import (
            "fmt"
            "time"
)

func PrintMessages(chVals chan int, done chan bool) {

    for {
        select {
        
        case v, ok := <- chVals:
            if !ok {
                done <- true      
                return
            }
            fmt.Println("Received", v)
        }
    }
}

func main() {
    ch := make(chan int)
    dn := make(chan bool)

    go PrintMessages(ch,dn)
    
    for i:=0 ; i < 10 ; i++ {
    
        ch <- i
    }
    close(ch)
    
    <-dn    // blocking until we receive something on this channel
    
    close(dn)
    fmt.Println("Exiting after all are done")
}