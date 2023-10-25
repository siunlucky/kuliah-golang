package main

import (
	"fmt"
);

func main()  {
	var x, y int;

	fmt.Scan(&x, &y)
	fmt.Println("");

	for i := 1; i <= x && i <= y; i++ {
		if x % i == 0 && y % i == 0 {
			fmt.Print(i, " ")
		}
	}
}