package main

import "fmt"

func main() {
    fullName := myfunction("carlos", "belisario")
    fmt.Println(fullName)

    anotherName, error := multiReturnFunction("john", "smith")
    if error != nil {
        fmt.Println("Handle Error Case")
    }
    fmt.Println(anotherName)

    fmt.Println("-------- anonymous func -------")
    myFunc := addOne()
    fmt.Println(myFunc()) // 2
    fmt.Println(myFunc()) // 3
    fmt.Println(myFunc()) // 4
    fmt.Println(myFunc()) // 5
}

/** function with args */
func myfunction(firstName string, lastName string) (string) {
  fullname := firstName + " " + lastName
  return fullname
}
/** multi return  */
func multiReturnFunction(firstName string, lastName string) (string, error) {
  return firstName + " " + lastName, nil
}

/** anonymous func */
func addOne() func() int {
  var x int
  // we define and return an
  // anonymous function which in turn
  // returns an integer value
  return func() int {
    // this anonymous function
    // has access to the x variable
    // defined in the parent function
    x++
    return x + 1
  }
}