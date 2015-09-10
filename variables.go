package main

import (
    "fmt"
)

// var statement declares a list of variables with the 
// type declared last.
// Since the variables are declared outside the main
// funcion, they are global and can be accessed anywhere
// in the package
// We can also declare variables of same type in the same
// line
var (
    name      string
    age       int
    location  string
)

func main(){

    name = "Mukul"
    age = 29
    location = "Bangalore"
    fmt.Println(name, age, location)


    // A var declaration can include intitializers one 
    // variable
    var (
        x int = 10
        y int = 20
    )

    fmt.Println(x, y)


    // Inferred typing
    // If an initializer is present the type can be omitted
    // the variable will take the tupe of the initializer
    var (
        n = "Prince Oberyn"
        a = 42
        l = "Dorne"
    )   

    fmt.Println(n, a, l)

    // variable can also be initialized on the same line
    var (
        e, f, g = 10, 20, 30
    )

    fmt.Println(e,f,g)

    // Inside a function the := short assignment can be used
    // in place of a var declaration with implicit type

    name1, location1 := "Prince Oberyn", "Dorne"
    age1 := 32
    fmt.Printf("%s (%d) of %s", name1, age1, location1)

    // A variable can contain any type, including functions
    // Outside a function, every construct begins with a
    // keyword (var, func and so on), and the := construct
    // is not available.

}
