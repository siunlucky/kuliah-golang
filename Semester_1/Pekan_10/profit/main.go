package main

import (
    "fmt"
)

func main() {
    var a, b, c float64;

    fmt.Scan(&a, &b)
	fmt.Println("");

	c = a - b
	if c > 0 {
		fmt.Println("Penurunan sebesar", c)
	} else if c == 0{
		fmt.Println("Tetap")
	} else {
		fmt.Println("Peningkatan sebesar", c * -1)
	}

}
