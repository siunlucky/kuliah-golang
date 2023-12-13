package main

import (
	"fmt"
)

func main() {
	var x string
	var hasil2 bool

	fmt.Scan(&x)

	// hasil1 = (65 <= x && x <= 90) || (97 <= x && x <= 122)
	hasil2 = ("a" <= x && x <= "z") || ("A" <= x && x <= "Z")

	// fmt.Println(hasil1);
	fmt.Println(hasil2);
}
 