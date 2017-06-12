package main

import (
        "fmt"
        "strconv"
)

func main() {
    // create an integer
    var a = 1
    // create a string
    var b = "this is a string" + strconv.Itoa(a)
    // output the length of the string
    fmt.Printf("This string is %d chars long [%s]\n", len(b), b)
    // for loop 1 to 10
    for i := 0 ; i < 10 ; i++ {
        fmt.Printf("%d\n", i + 1 )
    }
    
}