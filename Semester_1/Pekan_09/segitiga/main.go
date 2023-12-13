package main

import (
    "fmt"
)

func main() {
    var x, y, z int
    var hasil, hasil2, hasil3 bool

    fmt.Scan(&x, &y, &z)

    // Segitiga siku-siku
    hasil = ((x * x + y * y == z * z) || (x * x + z * z == y * y) || (y * y + z * z == x * x))

    // Segitiga sama sisi
    hasil2 = (x == y) && (y == z) && (x == z)

    // Segitiga Sama Kaki
    hasil3 = ((x == y) && x < z) || ((x == z) && z < y) || ((y == z) && z < x)


    fmt.Println("Segitiga siku-siku : ", hasil)
    fmt.Println("Segitiga sama sisi : ", hasil2)
    fmt.Println("Segitiga sama kaki : ", hasil3)
    fmt.Println("Bukan segitiga : ", !hasil3 && !hasil2 && !hasil)
}
