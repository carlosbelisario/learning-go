package main

import "fmt"

func main() {
   slices()
   maps()
   structs()
}

func arrays() [7]string {
   // declaring an array with elements
   days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
   fmt.Println(days[0])
   fmt.Println(days[3])
   return days
}

func slices() {
     days := arrays()
    /** slice declaration*/ 
    weekdays := days[0:5]
    fmt.Println(weekdays)
}

func maps() {
    youtubeSubscribers := map[string]int{
        "TutorialEdge":     2240,
        "MKBHD":            6580350,
        "Fun Fun Function": 171220,
    }

    fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350
}

func structs() {
    type Person struct {
        name string
        age  int
    }

    // our Team struct
    type Team struct {
        name    string
        players [2]Person
    }

    // declaring an empty 'Team'
    var myTeam Team
    fmt.Println(myTeam)

    players := [...]Person{Person{name: "Forrest"}, Person{name: "Gordon"}}
    // declaring a team with players
    celtic := Team{name: "Celtic FC", players: players}
    fmt.Println(celtic)
}
