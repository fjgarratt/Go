package main

import (
        "fmt"
)

func Concat( integers *[]int ) {
    *integers = append(*integers, 1)
    fmt.Println(integers)
}

func main() {
    tab := []int { -1, 0}
    Concat(&tab)
    fmt.Println(tab)
}