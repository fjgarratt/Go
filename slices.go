package main

import ("fmt")

func main() {
    
    vals := [3]int{0,2,4}
    
    s := vals[:]
    
    s[2] = 8
    s = append(s,128,4096)
    s[0] = 1
    fmt.Println(vals)
    fmt.Println(s)
}