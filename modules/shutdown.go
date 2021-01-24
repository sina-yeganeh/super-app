package main

import (
    "fmt"
    "log"
    "os/exec"
    "runtime"
)

// Dont have global variable !

func main() {
    // Checking the machine system
    if (runtime.GOOS == "linux") {
        LinuxShutdown()
    } else if (runtime.GOOS == "windows") {
        WindowsShutdown()
    } else {
        fmt.Println("Sorry, this tool cannot run in you machine!")
    }
}

func WindowsShutdown() {
    output, err := exec.Command("shutdown").Output()
    // Handel error
    if (err != nil) {
        log.Fatalln("Some thing went wrong:", err)
    }
    // Shut down system
    result := string(output)
    fmt.Println(result)
}

func LinuxShutdown() {
    output, err := exec.Command("shutdown -h Now").Output()
    if (err != nil) {
        log.Fatalln("Some thing went wrong:", err)
    }
    result := string(output)
    fmt.Println(result)
}
