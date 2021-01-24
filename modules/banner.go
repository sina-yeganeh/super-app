package main

import (
    "fmt"
    "ioutil"
    "log"
)

func ThreeDBanner() {
    banner, err := ioutil.ReadFile("3D.txt")
    if (err != nil) {
        log.Fatalln(err)
    }
    banner := string(banner)
    fmt.Println(banner)
}
