package main

import (
	"fmt"
)

func main() {
	var x, total int

	fmt.Scan(&x)
	fmt.Println("")

	for x > 0 {
        total += x % 10
        x /= 10
    }

	fmt.Println(total)

}
 