package main

import "fmt"

func main() {
    var x, n, gold, platinum, silver int;
	
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		x = 0
		for x < 50 {
			fmt.Scan(&x)
		}

		if x > 200 {
			gold++
		} else if x >= 100 && x <= 200{
			platinum++
		} else if x >= 50 && x < 100{
			silver++
		}
	}

	fmt.Print("Gold user : ", gold, ", Platinum user : ", platinum, ", Silver user : ", silver)
}
