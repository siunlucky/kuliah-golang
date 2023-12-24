package main

import (
    "fmt"
)

func main() {
    var a, b, c, x, y int;
	
	fmt.Scan(&a, &b, &c)
	fmt.Println("")

	if a > b {
        x = a
        y = b

        a = y
        b = x
    }

    if b > c {
        x = b
        y = c

        b = y
        c = x
    }

    if a > b { 
        x = a
        y = b

        a = y
        b = x
    }

	fmt.Println(a, b, c)
}
