package main

import (
    "fmt"
)

// A function can take zero or more arguments. The type comes after the variable
// name. Functions can be defined to return any number of values that are typed after.
func add(a, b int) int{
    return a + b
}

// In the following example, the location function returns two string values
func location(city string) (string, string) {
    var region string
    var continent string

    // We are switching on the variable city and the cases check for the value
    // of the city.
    switch city{
        case "Los Angeles", "LA", "Santa Monica":
            region, continent = "California", "North America"

        case "New York", "NYC":
            region, continent = "New York", "North America"

        default:
            region, continent = "Unknown", "Unknown"
    }
    return region, continent
}

// Functions take parameter. In Go, functions can return multiple "return parameter", not just
// a single value. They can be named and act just like variables.
// If the return parameters are named, a return statement without arguments returns the current 
// values of the results.
func location1(name, city string) (name1, continent string){
    switch city{

    case "New York", "LA", "Chicago":
        continent = "North America"
    default:
        continent = "Unknown"
    }

    name1 = name

    return 
}

func main(){
    fmt.Println(add(42, 13))

    region, continent := location("Santa Monica")

    fmt.Printf("Matt lives in %s, %s\n", region, continent)

    name, continent1 := location1("Matt", "LA")
    fmt.Printf("%s lives in %s\n", name, continent1)
}