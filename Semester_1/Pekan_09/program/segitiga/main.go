package main

import (
    "fmt"
)

func main() {
    var x, y, z int
    var hasil bool

    fmt.Scan(&x, &y, &z)

    hasil = x + y > z && x + z > y && y + z > x

    fmt.Println(hasil)
}