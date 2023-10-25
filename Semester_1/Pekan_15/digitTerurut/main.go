package main

import "fmt"

func main() {
	var x, prev, now int
	var isAscending, isDecreasing bool

	fmt.Scan(&x)

	prev = x % 10
	x /= 10

	isAscending, isDecreasing = true, true

	for x > 0 {
		now = x % 10
		if prev > now {
			isDecreasing = false
		} else if prev < now {
			isAscending = false
		}
		x /= 10

		prev = now
	}

	if isAscending {
		fmt.Println("ascending")
	} else if isDecreasing {
		fmt.Println("decreasing")
	} else {
		fmt.Println("tidak terurut")
	}
}
