package main

import "fmt"

func main() {
    fmt.Println("----------integers-----------")
    integer()
    fmt.Println("----------casting integers-----------")
    castingIntegers()
    fmt.Println("----------float-----------")
    float()
    fmt.Println("----------float casting integer-----------")
    converFloatAndInt()
    fmt.Println("----------boolean-----------")
    booleans()
    fmt.Println("----------string and const-----------")
    stringsAndConst()
}

/** review this function, why the result is -127*/
func integer() {
    var myint int8
    for i := 0; i < 129; i++ {
        myint += 1
    }
    fmt.Println(myint) // prints out -127
}

func castingIntegers() {
    var men uint8
    men = 5
    var women int16
    women = 6

    var people int
    // this throws a compile error
    /*people = men + women*/
    // this handles converting to a standard format
    // and is legal within your go programs
    people = int(men) + int(women)
    fmt.Println(people)
}


func float() {
    var maxFloat32 float32
    maxFloat32 = 16777216
    fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
    // it returns true
    fmt.Println(maxFloat32+10) // 16777216
    fmt.Println(maxFloat32+2000000) // 16777216
}

func converFloatAndInt() {
    // converting from int to float
    var myint int
    myfloat := float64(myint)
   
    // converting from float to int
    var myfloat2 float64
    myint2 := int(myfloat2)

    fmt.Print("integer casting to float: ")
    fmt.Println(myfloat)
    fmt.Print("float casting to integer: ")
    fmt.Println(myint2)
}

func booleans() {
    var isTrue bool = true
    var isFalse bool = false
    // AND
    if isTrue && isFalse {
      fmt.Println("Both Conditions need to be True")
    }
    // OR - not exclusive
    if isTrue || isFalse {
      fmt.Println("Only one condition needs to be True")
    }
}

func stringsAndConst() {
    var myName string
    myName = "Elliot Forbes"
    fmt.Println(myName)
    const meaningOfLife = 42
    fmt.Println(meaningOfLife)
}
