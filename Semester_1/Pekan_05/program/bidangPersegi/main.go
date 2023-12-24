package main

import (
	"fmt"
);

func main()  {
	var n, s, l, k int;

	fmt.Scan(&n);
	fmt.Println("");
	for i := 1; i <= n; i++ {
		fmt.Scan(&s);
		l = s * s
		k = s * 4
		fmt.Println(l, k);
	}
}