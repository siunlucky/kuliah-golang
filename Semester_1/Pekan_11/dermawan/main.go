package main

import "fmt"

func main() {
    var n, x, total int

	fmt.Scan(&n)
	total = 0
    for i := 0; i < n; i++ {
		fmt.Scan(&x)
		total += x
	}
	fmt.Println("")
	fmt.Println(total)
}
