package main

import (
        "fmt"
)

const (
    _ = iota
    START_PORT = iota << ( 10 * iota )
    MID_PORT
    END_PORT 
)

func main() {
    fmt.Println(START_PORT)
    fmt.Println(MID_PORT)
    fmt.Println(END_PORT)
}