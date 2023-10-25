package main

import (
	"fmt"
);

func main()  {
	var x, count int;
	var isPrima bool;

	fmt.Scan(&x)
	fmt.Println("");

	count = 0
	isPrima = true;

	for i := 1; i <= x; i++ {
		if x % i == 0 {
			count++
		}
	}

	if count != 2 {
		isPrima = false
	}

	fmt.Println(isPrima)
}