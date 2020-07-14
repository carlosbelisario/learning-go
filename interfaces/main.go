package main

import (
    "fmt"
    "time"
    "math"
)

/** when the param type is an interface we can pass anything we want*/
func myFunction(param interface{}) {
    fmt.Println(param)
}

func main() {
    var my_age int
    my_age = 33
    myFunction(my_age)
    name := "Carlos"
    myFunction(name)

    // testing custom interface
    var baseGuitarist BaseGuitarist
    baseGuitarist.Name = "Carlos"
    baseGuitarist.PlayGuitar()
    var acousticGuitarist AcousticGuitarist
    acousticGuitarist.Name = "Carlitos"
    acousticGuitarist.PlayGuitar()

    var guitarists []Guitarist
    guitarists = append(guitarists, baseGuitarist)
    guitarists = append(guitarists, acousticGuitarist)
    fmt.Println(guitarists)

    // interface with return
    var programmers []Employee
    carlos := Engineer{Name: "Carlos", Birthday: time.Date(1986, 12, 20, 0, 0, 0, 0, time.UTC)}
    programmers = append(programmers, carlos)
    fmt.Println(programmers)
    fmt.Println(programmers[0].Age())
}

/* custom interface */

type Guitarist interface {
    PlayGuitar()
}

type BaseGuitarist struct {
    Name string
}

type AcousticGuitarist struct {
    Name string
}

func (guitarist BaseGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the base guitar\n", guitarist.Name)
}

func (guitarist AcousticGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the acoustic guitar\n", guitarist.Name)
}

// interface with return value

type Employee interface {
    Language() string
    Age() int
    Random() (string, error)
}

type Engineer struct {
    Name string
    Birthday time.Time
}

func (e Engineer) Language() string {
    return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
     today := time.Now()
    return int(math.Floor(today.Sub(e.Birthday).Hours() / 24 / 365))
}

func (e Engineer) Random() (string, error) {
    return "this is a random string", nil
}