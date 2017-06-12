package main

import (
    "fmt"
    "time"
)
var chReply = make(chan string)

func ContactMars() {
    time.Sleep(10 * time.Second)
    chReply <- "We're coming!"
}

func main() {
    go ContactMars()
    
    chTimer := time.After(3*time.Second)
    
    select {
    case msg := <-chReply:
        fmt.Println(msg)
    case <-chTimer:
        fmt.Println("Timeout")
    }
    
    fmt.Println("Done")
}