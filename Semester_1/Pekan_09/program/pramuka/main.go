package main

import (
    "fmt"
)

func main() {
    var n, i int;
    var a, b, c, d, e, hasil bool;

	fmt.Scan(&n)

	hasil = true;
    i = 0
    for i < n && hasil == true {
        fmt.Scan(&a, &b, &c, &d, &e)
        i++
        hasil = a == true && b == true && c == true && d == true && e == true
    }

	fmt.Println("")
    fmt.Println(hasil)
}
