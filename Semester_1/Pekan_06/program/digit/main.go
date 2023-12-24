package main

import (
	"fmt"
);

func main()  {
	var x, y, total int;

	fmt.Scan(&x);

	total = 0;
	for x > 0 {
		y = x % 10
		x = x / 10
		fmt.Print(y, " ");
		total += y;
	 }
	
	fmt.Println("");
	fmt.Println(total);
}