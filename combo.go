package main

import (
        "fmt"
        "strings"
)

var a = "ABCD"

func main() {
        
    b := strings.Split(a, "")
    n := len(b)
    
    permute(b, 0, n-1)
}

func swap( a[]string, first int, second int ) {
    
        temp := a[second]
        a[second] = a[first]
        a[first]  = temp
}

func permute( a[]string, l int, r int) {
    
    if ( l == r ) {
        fmt.Printf("%s\n", a)
    } else {
     
         for i := l ; i <= r ; i++ {
         
             swap(a, l, i )
             permute( a, l+1, r )
             swap(a, l, i )
         }
    }
}


