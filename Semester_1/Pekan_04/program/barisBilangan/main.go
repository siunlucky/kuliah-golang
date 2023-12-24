package main

import "fmt";

func main() {
	var a, b, c, d, e int
	
	fmt.Scan(&a, &b)
	c = a + b
	d = b + c
	e = c + d
	fmt.Println(c, d, e)
}