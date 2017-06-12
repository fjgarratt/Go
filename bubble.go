package main

import (
        "fmt"
)

var a = []int {102,5,3,76,34,21,67,43,31,0,56,98}

func main() {

    fmt.Println( bubble_sort( a ) )
}

func bubble_sort( c []int ) []int {
    
    var swap, pos int = 0, 0
    
    if len(c) > 1 {
        
        for {          
            if pos+1 >= len(c) {
                pos = 0
                
                if swap == 0 {
                    break
                }
                swap = 0
            }
                    
            if c[pos] > c[pos+1] {
                
                c = swapit( c, pos)
                fmt.Println( c )
                swap = 1
            }
            
            pos++
        }
    }
    
    return c
}

func swapit( c []int, pos int ) []int {
    
    var tmp int = c[pos]
    
    c[pos] = c[pos+1]
    c[pos+1] = tmp
    
    return c
}