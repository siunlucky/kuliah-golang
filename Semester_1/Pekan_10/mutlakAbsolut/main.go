package main

import (
    "fmt"
)

func main() {
    var x int;

    fmt.Scan(&x)
	fmt.Println("");

	if x < 0 {
        x = x * -1
    }

	fmt.Println(x)
}
