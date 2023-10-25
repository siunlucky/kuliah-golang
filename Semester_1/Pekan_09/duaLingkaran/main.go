package main

import (
    "fmt"
)

func main() {
    var x, y, z int
    var hasil bool

    fmt.Scan(&x, &y, &z)
	fmt.Println("");

    hasil = x + y < z

    fmt.Println(hasil)
}
