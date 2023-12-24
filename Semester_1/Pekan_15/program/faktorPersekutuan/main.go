package main

import (
	"fmt"
);

func main()  {
	var x, y, i int;

	fmt.Scan(&x, &y)
	fmt.Println("");

	i = 1
	for i <= x && i <= y {
		if x % i == 0 && y % i == 0 {
			fmt.Print(i, " ")
		}
		i++
	}
}