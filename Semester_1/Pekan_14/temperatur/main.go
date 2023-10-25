package main

import (
	"fmt"
);

func main()  {
	var x, min, max, lastX, sum, n int;
	var average float64;
	var isLoop bool;

	isLoop = true
	average, min, max, lastX, sum, n = 0, 0, 0, 0, 0, 0;

	for isLoop == true {
		fmt.Scan(&x)
		if lastX == 0 && x == 0 {
			isLoop = false;
		}

		if min > x {
			min = x
		}

		if max < x {
			max = x
		}

		sum += x
		lastX = x
		n++
	}

	average = float64(sum) / float64(n-1)
	
	fmt.Println("")
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(average)
}