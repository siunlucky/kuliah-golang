package main

import (
	"fmt"
)

func main() {
	var x string
	var hasil bool

	fmt.Scan(&x)

	hasil = ("a" <= x && x <= "z") || ("A" <= x && x <= "Z")

	fmt.Println(hasil);
}
 