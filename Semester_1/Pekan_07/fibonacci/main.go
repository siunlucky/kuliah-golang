package main

import (
	"fmt"
)

func main() {
	var n int
	var a, b int;
	
    
    fmt.Scan(&n)
	fmt.Println("")

    a = 0;
	b = 1;
    for i := 0; i <= n; i++ {
        fmt.Print(a, " ")

		a, b = b, a + b
		// Code diatas dapat ditulis seperti ini
		// x := a
		// y := b

        // a = y
		// b = x + y
    }
}
 