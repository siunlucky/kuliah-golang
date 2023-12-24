package main

import (
	"fmt"
)

func main() {
	var a, b, x, y, n, i int;
	
    
    fmt.Scan(&n)
	fmt.Println("")

    a = 0;
	b = 1;
    for i = 0; i <= n; i++ {
        fmt.Print(a, " ")

		x = a
		y = b

        a = y
		b = x + y

		// Code diatas dapat ditulis seperti ini
		// a, b = b, a + b
    }
}
 