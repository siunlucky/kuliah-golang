package main

import (
	"fmt"
);

func main()  {
	var x, hasil, digit int;

	fmt.Scan(&x)
	fmt.Println("");

	hasil = x % 10
	for x > 0 {
		digit = x % 10
		if digit > hasil {
			hasil = digit
		}
		x /= 10
	}

	fmt.Println(hasil)
}