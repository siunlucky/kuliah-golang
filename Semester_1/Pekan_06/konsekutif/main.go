package main

import (
	"fmt"
);

func main()  {
	var x, lastDigit, digitNow int;
	var isOne bool;

	fmt.Scan(&x); // 12345

	isOne = true;
	
	for x > 0 && isOne == true {
		lastDigit = x % 10; // 12345 % 10 = 5
		x /= 10; // 12345 / 10 = 1234
		digitNow = x % 10 // 1234 % 10 = 4
		isOne = (lastDigit - digitNow == 1) || (digitNow - lastDigit == 1); // (5 - 4 == 1) || (4 - 5 == 1) = True
		x /= 10 // 1234 / 10 = 123
	}

	fmt.Println(isOne);
}