package main 

import ( 
	"fmt"
)

func main() {

	// fibonacci series
	var a [10]int	// number of places

	last, current:= 0, 1
	
	a[0] = 1
	for i:= 1; i< 10 ; i++ {
		a[i] = last + current
		
		last = current
		current = a[i]
	}

	fmt.Println(a)
}



