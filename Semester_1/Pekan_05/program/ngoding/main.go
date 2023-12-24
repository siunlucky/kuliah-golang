package main

import (
	"fmt"
);

func main()  {
	var n, x, sum int
	var hasil float64;

	fmt.Scan(&n)

	sum = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)

		sum += x
	}

	hasil = float64(sum) / float64(n)

	fmt.Printf("%.3f" , hasil)

}