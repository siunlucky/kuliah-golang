package main

import (
    "fmt"
)

func main() {
    var x, y, z float64
    var hasil bool

    fmt.Scan(&x, &y, &z)
	fmt.Println("");

    hasil = (x + y) / 2.0 == z || (x + z) / 2.0 == y || (y + z) / 2.0 == x

    fmt.Println(hasil)
}
