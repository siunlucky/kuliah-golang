package main

import (
    "fmt"
)

func main() {
    var a, b, c, d, e string;

	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Println("")

	if a == "kalah" && b == "kalah" && c == "kalah" && d == "kalah" && e == "kalah" {
		fmt.Println("dipecat")
	} else {
		fmt.Println("tidak dipecat")
	}
}
