package main

import "fmt"

func main() {
    var x int;

	fmt.Scan(&x)

	if x == 0 {
		x = 50
	}
	
	if x > 200 {
		fmt.Println("Gold user")
	} else if x >= 100 && x <= 200{
		fmt.Println("Platinum user")
	} else if x >= 50 && x < 100{
		fmt.Println("Silver user")
	}
}
