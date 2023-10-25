package main

import (
    "fmt"
)

func main() {
    var x, y, z int
    var hasil bool

    fmt.Scan(&x, &y, &z)
	fmt.Println("");

    hasil = float64(x + y) / 2.0 == float64(z) || float64(x + z) / 2.0 == float64(y) || float64(y + z) / 2.0 == float64(x)

    fmt.Println(hasil)
}
