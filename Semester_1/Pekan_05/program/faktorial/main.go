package main

import (
	"fmt"
);

func main()  {
	var n int

	fmt.Scan(&n)

	for i := n-1; i >= 1; i-- {
		n *= i 
	}

	fmt.Println(n);

}