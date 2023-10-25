package main

import (
    "fmt"
)

func main() {
    var a, b, c int;
	
	fmt.Scan(&a, &b, &c) // 5 8 3
	fmt.Println("")

	// a = 5, b = 8, c = 3

	if a > b { // 5 > 8
        a, b = b, a 
    }

	// a = 5, b = 8, c = 3

    if b > c { // 8 > 3
        b, c = c, b // b = 3, c = 8
    }

	// a = 5, b = 3, c = 8

    if a > b { // 5 > 3
        a, b = b, a // a = 3, b = 5
    }

	// a = 3, b = 5, c = 8

	fmt.Println(a, b, c)
}
