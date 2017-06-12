package main

import (
            "fmt"
)

func PrintMessages(chVals chan int, done chan bool) {
    for x := range chVals{
        fmt.Println(x)
    }
    done <- true
}

func main() {
    ch := make(chan int)
    dn := make(chan bool)

    go PrintMessages(ch,dn)
    /*
    go func() {
    
        for x := range ch {
            fmt.Println(x)
        }
    }()*/
    
    for i:=0 ; i < 10 ; i++ {
    
        ch <- i
    }
    close(ch)
    
    <-dn    // blocking until we receive something on this channel
    
    close(dn)
    fmt.Println("Exiting after all are done")
}