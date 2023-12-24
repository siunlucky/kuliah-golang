package main

import (
    "fmt"
)


func main() {
    var x, binary, remainder, counter int

    fmt.Scan(&x)
    binary = 0
    counter = 1
    remainder = 0

    for x != 0 {
        remainder = x % 2
        x = x / 2
        binary += remainder * counter
        counter *= 10
    }

    fmt.Println(binary)

}