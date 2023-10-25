package main

import "fmt"

func main() {
    var x, y int
    var binaryString string

    fmt.Scan(&x)
    fmt.Println("")

    for x > 0  {
        y = x % 2
        x = x / 2
        binaryString = fmt.Sprint(y) + binaryString
    }

    fmt.Println(binaryString)
}
