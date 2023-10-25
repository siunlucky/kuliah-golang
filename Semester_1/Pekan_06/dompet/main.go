package main

import (
	"fmt"
);

func main()  {
	var x, total int;

	total = 0;
	x = 1;

	for x != 0 {
		fmt.Scan(&x)
		total += x
	}

	fmt.Println(total);

}