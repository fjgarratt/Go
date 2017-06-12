package main

import (
        "fmt"
        "math/rand"
        "time"
        "unicode"
)

var suit = map[string]string {"clubs":"black","spades":"black","hearts":"red","diamonds":"red"}
var card = []string {"ace","two","three","four","five","six","seven","eight","nine","ten","jack","queen","king"}

type suits struct {
    name   string
    colour string
    cards  []string
}

type deck struct {
    pack []suits
}

func Create_deck() deck {
        
    my_deck := deck{}
    
    for k,v := range suit {
        
        this_suit := suits{k,v,card}
        
        my_deck.pack = append(my_deck.pack, this_suit)
    }
    
    return my_deck
}

func Shuffle_deck(d deck) []string {
    
    var shuffled []string
    
    done := make(map[int]int)
    
    for {
    
        r := get_rand(0,len(d.pack))
        if len(d.pack[r].cards) == 0 {
            done[r] = 1
            
            if len(done) == 4 {
                break
            }
            continue
        }
                
        rnumb := get_rand(0,len(d.pack[r].cards))
        
        shuffled        = append(shuffled, fmt.Sprintf("%s of %s", upper(d.pack[r].cards[rnumb]), upper(d.pack[r].name) ))
        d.pack[r].cards = remove(d.pack[r].cards, rnumb)
    }

    return shuffled
}

func get_rand(min, max int) int {
    
    return rand.Intn(max - min) + min
}

func remove(s []string, i int) []string {
    
    s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-1]
}

func upper(s string) string {
    
    a := []rune(s)
    a[0] = unicode.ToUpper(a[0])
    
    s = string(a)
    
    return s
}

func main() {
    rand.Seed(time.Now().Unix())
    
    my_deck := Create_deck()
    shuffle := Shuffle_deck( my_deck )
    
    for i, card := range shuffle {
    
        fmt.Printf("%d\t%s\n", i+1, card)
    }
}