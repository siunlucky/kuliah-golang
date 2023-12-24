package main

import "fmt"

func main(){
	var x, a, b, c, d, e int

	fmt.Scan(&x) // 12345

	a = x / 10000 // 1
	x = x % 10000 // 2345
	b = (x / 1000) * 10 // 20
	x = x % 1000 // 345
	c = (x / 100) * 100 // 300
	x = x % 100 // 45
	d = (x / 10) * 1000 // 4000
	e = (x % 10) * 10000 // 50000
 
	fmt.Println(a + b + c + d + e)
}