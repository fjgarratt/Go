package main

import (
        "fmt"
)

var a = []string {"hayley","fenton","darcey","kingsley","webster","holly","sophie","lucy","lynsey","chrystel","gary","daisy"}

func main() {

    fmt.Println( bubble_sort( a ) )
}

func bubble_sort( c []string ) []string {
    
    var swap, pos, lateral int = 0, 0, 0
    
    if len(c) > 1 {
        
        for {
            if pos+1 >= len(c) {
                pos = 0
                
                if swap == 0 {
                    break
                }
                swap = 0
            }

            if lateral > (len(c[pos])-1) {

                lateral = 0
                pos++       
                continue
            }
            
            var a = []rune(c[pos])[lateral]
            var b = []rune(c[pos+1])[lateral]
            
            if a > b {
                
                c = swapit( c, pos)
                swap = 1

            } else if a == b {
    
                lateral++                            
                continue
            }
            
            lateral = 0
            pos++
        }
    }
    
    return c
}

func swapit( c []string, pos int ) []string {
    
    var tmp string = c[pos]
    
    c[pos] = c[pos+1]
    c[pos+1] = tmp
    
    return c
}