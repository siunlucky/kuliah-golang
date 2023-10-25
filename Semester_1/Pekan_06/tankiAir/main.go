package main

import (
	"fmt"
);

func main()  {
	var t, v, tv int;
	var isFull bool;

	fmt.Scan(&t);

	tv = 0;
	for tv < t {
		fmt.Scan(&v);
		tv += v;
		isFull = tv >= t;
		fmt.Println(isFull)
	}
}