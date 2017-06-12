package main

import (
            "fmt"
)

//var NumQueries int

func MakeConnectedWrapper(userFunc func() bool) func() bool {
  
    var NumQueries int
    
    return func() bool {
        
        result := Connected()
        NumQueries++
        fmt.Println("NumQueries: ", NumQueries)   
        return result
    }
}

func Connected() bool {
    // ...
    return true
}

func main() {
    //fmt.Println(NumQueries)
    //Connected()
    connected := MakeConnectedWrapper(Connected)
    connected()
    connected()
    connected()
    //fmt.Println(NumQueries)
}
