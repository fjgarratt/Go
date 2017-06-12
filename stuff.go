package main

import (
        "fmt"
)

type stuff struct {
            new_videos int
}

type m map[string]*stuff

func main() {
    // don't know how to assign values to this construct
    // check this out
    //k := map[string]stuff {}

    p := make(m)
     
    p["pluto"] = &stuff{5}
   
    p["pluto"].new_videos++
    
    fmt.Println(p["pluto"])

}

