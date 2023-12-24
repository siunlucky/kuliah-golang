package main

import (
	"fmt"
);

func main()  {
	var x, lastDigit, digitNow int;
	var isOne bool;

	fmt.Scan(&x);

	isOne = true;
	
	for x > 0 && isOne == true {
		lastDigit = x % 10;
		x /= 10;
		digitNow = x % 10
		isOne = (lastDigit - digitNow == 1) || (digitNow - lastDigit == 1);
		x /= 10
	}

	fmt.Println(isOne);
}