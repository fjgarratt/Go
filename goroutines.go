package main

import ("fmt")

func PrintMessage(m string) {
    fmt.Println(m)
    Done = true
}

var Done bool

func main() {
    go PrintMessage("Hello")
    
    for {
        if Done {
    
            fmt.Println("Goroutines")
            return
        }
    }
}