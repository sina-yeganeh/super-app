package main

import (
    "fmt"
    "math/rand"
)

// Global variable
var (
    userNumber int
    computerNumber int
    peroid int
)

func main() {

}

func GetNumberAndPeroid() {
    fmt.Println("Plaese enter a number:[0, 200] ")
    fmt.Scan(&userNumber)
    fmt.Println("Plaese enter a peroid: ")
    fmt.Scan(&peroid)
}

func GuessNumber() {
    computerNumber = rand.Intn()
}
