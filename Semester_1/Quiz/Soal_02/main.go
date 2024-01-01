package main

import "fmt"

func main() {
	var sum int
	var isSuwir, isCakue, isAti, isTelur bool

	fmt.Scan(&isSuwir, &isCakue, &isAti, &isTelur)

	sum = 6000
	if isSuwir {
		sum += 3000
	}

	if isCakue {
		sum += 1500
	}

	if isAti {
		sum += 4500
	}

	if isTelur {
		sum += 4000
	}

	fmt.Println(sum)

}
