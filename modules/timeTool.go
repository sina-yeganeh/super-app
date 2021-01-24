package main

import (
    "fmt"
    "time"
)

var (
    property = [1] string {
        "Time",
    }
    number int
)

func main() {
    fmt.Println("[1]", property[0])
    fmt.Println(">>> Plaese chose the number: ")
    fmt.Scan(&number)
    for i := 0; ; i++ {
        switch (number) {
        case 1:
            Time()
        default:
            fmt.Println("Please chose the number :(")
        }
    }
}

func Time() {
    timeNow := time.Now()
    timeNow = string(timeNow)
    fmt.Println(timeNow)
}
