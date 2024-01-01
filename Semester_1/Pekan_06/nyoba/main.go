package main

import "fmt"

func main() {
	var a, n, b, un, sn int

	a = 4
	un = 1000000001
	b = 4

	n = (un - a + b) / b
	sn = n / 2 * (un + a)

	fmt.Println(n)
	fmt.Println(sn)

	// a = suku pertama
	// un = suku terbesar
	// b = selisih

	// var x, y, total int

	// x = 125524238436
	// y = 4
	// total = 0

	// for y != x {
	// 	y = y * 3
	// 	total = total + y
	// }
	// fmt.Println(total+4, "")

}
