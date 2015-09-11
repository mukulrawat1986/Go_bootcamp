package main

import (
    "fmt"
)

// Go has pointers, but no pointer arithmetic. Struct fields can be accessed through a 
// struct pointer. The indirection through the pointer is transparent (you can directly
// call fields and methods on a pointer.)

// By default Go passes arguments by values (copying the argument), if you want to pass the
// argument by reference you need to pass pointers (or use a structure using reference values like
// slice or maps)

// To get the pointer of a value, use the & symbol in front of the value, to dereference a pointer
// use the * symbol
func zeroval(ival int){
    ival = 0
}

func zeroptr(iptr *int){
    *iptr = 0
}


func main(){
    i := 1

    fmt.Println("initial: ", i)

    zeroval(i)
    fmt.Println("zeroval: ", i)

    zeroptr(&i)
    fmt.Println("zeroptr: ", i)

    fmt.Println("pointer: ", &i)
}