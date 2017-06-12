package main

import (
	"fmt"
	"math"
	"strings"
)

func makeLines() [100][]rune {
	var lines [100][]rune
	for i := 0; i < 100; i++ {
		l := int(math.Abs(math.Sin(0.05*float64(i)*3.14) * 100.0))
		str := strings.Repeat("*", int(l))
		for c := 0; c < l; c++ {
			lines[i] = append(lines[i], rune(str[c]))
		}
	}
	return lines
}

func main() {

	lines := makeLines()

    for _, l := range lines {
		
		for _, k := range l { 
         
		    fmt.Printf("%c", k)
		}
		fmt.Println("\n")
    }

	fmt.Println("This is a base code for an exercice, this file is not meant to be run \"as is\"\n")
	lines = lines
}

// TIP : use %c to print a character (rune) with fmt.Printf