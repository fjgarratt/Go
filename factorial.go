package main 

import ( 
	"fmt"
)

func main() {

	num := readline()
	run := num
	for {

		if num - 1 > 0 {
			fmt.Printf("Multiplying %d by %d (%d)\n", num, (num-1), run)
			run = run * ( num - 1 )
		} else {

			break
		}

		num = num - 1
	}

	fmt.Println( run )
}



