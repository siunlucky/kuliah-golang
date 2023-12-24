package main

import (
	"fmt"
)

func main() {
	var x, a, b, c, d int
	var diskon, cashback bool

	fmt.Scan(&x)
	fmt.Println("")

	a = x / 1000
	b = (x / 100) % 10
	c = (x / 10) % 10
	d = x % 10

	diskon = ((b * 10 + c) % 2 == 0)
	cashback = d != 0 && (a + c) % d == 0

	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
}
 