package main

import (
    "fmt"
    "math/rand"
    "os/exec"
)

// Global variable
var (
    userAnswer int
    computerAnswer int
    computerResult int
)

func main() {
    // Roor form
    for i := 0; ; i++ {
        fmt.Pribtln("[0] or [exit] exit")
        fmt.Println("[1] Rock")
        fmt.Println("[2] Cuter")
        fmt.Println("[3] Paper")
        fmt.Println(">>> Please enter the number: ")
        fmt.Scan(&userAnswer)
        // Check to user want exit from tool
        if (userAnswer == 0) || (userAnswer == "exit") {
            output, err := exec.Command("^C").Output()
            if (err != nil) {
                fmt.Println(err)
            }
            result := string(output)
            fmt.Println(result)
        // If all think is true, call ComputerAnswer function
        } else {
            computerResult := ComputerAnswer()
            CheckResult()
        }
    }
}

func ComputerAnswer() int {
    computerAnswer := rand.Intn(3)
    // Check if computer answer == 0 and then
    // make another number
    if (computerAnswer == 0) {
        for i := 0; ; i++ {
            computerAnswer := rnad.Intn(3)
            if (computerAnswer == 0) {
                continue
            } else {
                // If computer answer is true then
                // return that number
                return computerAnswer
            }
        }
    } else {
        return computerAnswer
    }
}

func CheckResult() {
    // Chose loser or winer with if, else if stament
    if (userAnswer == 1) && (computerResult == 3) {
        fmt.Println("You win!")
    } else if (userAnswer == 1) && (computerResult == 2) {
        fmt.Println("Computer win!")
    } else if (userAnswer == 2) && (computerResult == 3) {
        fmt.Println("You win!")
    } else if (userAnswer == 2) && (computerResult == 1) {
        fmt.Println("Computer win!")
    } else if (userAnswer == 3) && (computerResult == 1) {
        fmt.Println("You win!")
    } else if (userAnswer == 3) && (computerResult == 2) {
        fmt.Println("Computer win!")
    } else {
        fmt.Println("It is bug!!! Please report it!")
    }
}
