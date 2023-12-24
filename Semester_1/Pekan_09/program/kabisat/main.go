package main

import (
    "fmt"
)

func main() {
    var x int
    var hasil bool

    fmt.Scan(&x)
	fmt.Println("");

    hasil = (x % 4 == 0 && x % 100 != 0) || x % 400 == 0
	// Hasil akan true jika (x dapat dibagi habis 4 AND x tidak dapat dibagi habis 100) OR x dapat dibagi habis 400

    fmt.Println(hasil)
}
