package main

import (
	"fmt"
);

func main()  {
	var x, y, digit int;
	var hasil bool;

	fmt.Scan(&x, &y)
	fmt.Println("");

	hasil = false
	for y > 0 {
		digit = y % 10
		if digit == x {
			hasil = true
		}
		y /= 10
	}

	fmt.Println(hasil)
}