package main

import (
        "fmt"
        "os"
        )

func OpenCloseDoor() {

    // this defer func is executed last before exiting the OpenCloseDoor function
    // executed in LIFO
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recovering ", err)
            os.Exit(1)
        }
        fmt.Println("Door closing")
    }()

    panic("A stone is blocking the door")
    
    fmt.Println("Door Opened") 
}

func main() {

    OpenCloseDoor()
}