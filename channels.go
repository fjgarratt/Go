package main

import (
            "fmt"
            "time"
)

func myFunc(ch chan int) {

    x := <- ch
    fmt.Println(x)
}

func main() {

    ch := make(chan int)
    ch <- 2016
    go myFunc(ch)
    
    time.Sleep(1*time.Second)
}