package main 

import ( 
	"fmt"
)

func main() {

	name := "fenton"
	returnOfFunc(&name)

    fmt.Println("Hello from", name )
}

func returnOfFunc(name *string) {

	*name = "booyaa"
}