package main

import (
    "fmt"
    "runtime"
    "os/exec"
    "scolor"
)

func main() {
    // Check the machine's system
    if (runtime.GOOS == "linux") {
        output, err := exec.Command("clear").Output()
        if (err != nil) {
            fmt.Println(err)
        }
        result := string(output)
        fmt.Println(result)
    } else if (runtime.GOOS == "windows") {
        output, err := exec.Command("clear").Output()
        if (err != nil) {
            fmt.Println(err)
        }
        result := string(output)
        fmt.Println(result)
    } else { // If machine not Windows or Linux
        fmt.Println("Sorry, We can't run in you machine!")
    }
    // Import the function
    CalculatorRoot()
}

func CalculatorRoot() {
    var (
        operator = [] string {
            "Total",
            "Minus",
            "Beat",
            "Division",
        }
    )
    // scolor.SetTextColor("Green", "[1] " + operator[0])
}
