package main

import "fmt"

func main() {
    var x, sum, gold, platinum, silver int;

	sum = 0
	for sum < 500 {
		x = 0
		for x < 50 {
			fmt.Scan(&x)
		}

		sum += x

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
