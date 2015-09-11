package main

import (
    "fmt"
)

// After importing a package, you can refer to the name it exports(meaning variables, methods
// and functions that are available from the package). In Go, a name is exported if it begins 
// with a capital letter.


// Constants can be declared like variable, but with the const keyword
// Constants can only be characters, string, boolean or numeric values 
// and cannot be used using the := syntax.
// An untyped constant takes the type needed by its constant

const (
    Pi = 3.14
    Truth = false
    Big = 100
    small = 1
)

func main(){
    fmt.Println(Pi)
    fmt.Println(Truth)
    fmt.Println(Big)
    fmt.Println(small)

    // We use fmt.Printf is used when you want to print one or multiple values using
    // a defined format specifier
    // fmt.Sprintf formats according to a format specifier and returns the resulting string
    aka := fmt.Sprintf("Number %d", 6)

    fmt.Printf("%s\n", aka)
}