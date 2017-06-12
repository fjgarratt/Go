package main

import (
        "fmt"
        "time"
        "sync"
)

func DoThis(w *sync.WaitGroup) {
    
    fmt.Println(time.Now())
    w.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(10)
    
    for i:=0; i<10; i++ {
        go DoThis(&wg)
    }
    
    wg.Wait()
    fmt.Println("Here finished")
}