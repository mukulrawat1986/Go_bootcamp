package main

import (
    "fmt"
)

// In Go only constants are immutable
// However, because arguments are passed by value, a function
// receiving a value argument and mutating it won't mutate the original value

// we define a structure here
type Artist struct {
    Name, Genre string
    Songs int
}

func newRelease(a Artist) int{
    a.Songs ++
    return a.Songs
}

// To mutate the passed value we need to pass it by reference using a pointer
func newRelease_mutate(a *Artist) int{
    a.Songs ++
    return a.Songs
}


func main(){
    me := Artist{Name: "Matt", Genre: "Electric", Songs: 42}
    fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
    fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)

    me1 := &Artist{Name: "Matt", Genre: "Electric", Songs: 42}
    fmt.Printf("%s released their %dth song\n", me1.Name, newRelease_mutate(me1))
    fmt.Printf("%s has a total of %d songs\n", me1.Name, me1.Songs)
}