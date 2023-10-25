package main

import (
    "fmt"
)

func main() {
    var n int;
    var a, b, c, d, e, hasil bool;

	fmt.Scan(&n)

	hasil = true;

    for i := 0; i < n && hasil == true; i++ {
        fmt.Scan(&a, &b, &c, &d, &e)
        hasil = a == true && b == true && c == true && d == true && e == true
    }

	fmt.Println("")
    fmt.Println(hasil)
}
