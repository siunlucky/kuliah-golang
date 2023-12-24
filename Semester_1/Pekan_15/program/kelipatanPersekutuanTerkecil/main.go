package main

import "fmt"

func main() {
    var x, y, a, b, k, l, kpk int

    fmt.Scan(&x, &y)
	fmt.Println("")

	// a, b = x, y
    a = x
    b = y

	for b != 0 {
        // a, b = b, a % b
        k = b
        l = a % b

        a = k
        b = l
    }

    kpk = (x * y) / a

    fmt.Println(kpk)
}
