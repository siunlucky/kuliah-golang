package main

import (
	"fmt"
);

func main()  {
	var n int;
	var x string;

	fmt.Scan(&n, &x);
	fmt.Println("");
	for i := 1; i <= n; i++ {
		fmt.Println(x);
	}
}