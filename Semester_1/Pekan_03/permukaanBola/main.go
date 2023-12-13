package main

import "fmt";

func main() {
	var r, pi, lpb float64
	fmt.Scan(&r)

	pi = 22.0 / 7.0

	lpb = 4.0 * pi * (r * r)
	fmt.Println(lpb)
}