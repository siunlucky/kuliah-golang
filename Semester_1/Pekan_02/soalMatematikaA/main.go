package main

import "fmt";

func main() {
	var x, y int
	var f float64

	fmt.Scan(&x, &y)

	f = (1 / ((3 * (float64(x) * float64(x))) + 10)) + (10 * float64(y)) + 7
	fmt.Println(f);
}