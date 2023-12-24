package main

import (
	"fmt"
);

func main()  {
	var n, x, hasil int;

	fmt.Scan(&n)
	fmt.Println("");

	hasil = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if hasil < x {
			hasil = x
		}
	}

	fmt.Println(hasil)
}