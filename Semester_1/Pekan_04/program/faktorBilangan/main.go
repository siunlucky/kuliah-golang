package main

import (
    "fmt"
)

func main() {
	var x, y int
	var isFactor bool

	fmt.Scan(&x, &y)

	isFactor = y % x == 0

	fmt.Println(isFactor)
}