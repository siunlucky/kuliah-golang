package main

import (
	"fmt"
);

func main()  {
	var x int;

	fmt.Scan(&x)
	fmt.Println("")

	for i := 1; i <= x; i++ {
		for j := 1; j <= x; j++ {
			if i == j || j == x - i + 1 {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}