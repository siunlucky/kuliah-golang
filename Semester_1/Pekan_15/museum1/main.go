package main

import "fmt"

func main() {
    var x, increase, lastVisitor, sum, n int;
	var average float64;

	lastVisitor, sum, increase, n = 0, 0, 0, 0

	for x >= 0 && x <= 200 {
		fmt.Scan(&x)

		if x >= 0 && x <= 200 {
			if x > lastVisitor {
				increase++
			}
			sum += x
			n++
		}
	}

	average = float64(sum) / float64(n)

	if average != average {
        average = 0

		fmt.Print(increase, average)
    } else {
		fmt.Printf("%d %.2f", increase, average)
	}
	
    
}
