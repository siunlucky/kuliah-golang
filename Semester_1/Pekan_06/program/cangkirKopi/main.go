package main

import (
	"fmt"
);

func main()  {
	var n, m, x, y, total int;

	fmt.Scan(&n, &m, &x, &y);

	total = 0;
	for x <= n && y <= m {
		n -= x
		m -= y;
		total++
	}
	fmt.Println(total)
}