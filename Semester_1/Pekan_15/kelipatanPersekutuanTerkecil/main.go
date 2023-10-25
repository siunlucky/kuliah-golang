package main

import "fmt"

func main() {
    var x, y, a, b int

    fmt.Scan(&x, &y)
	fmt.Println("")

	a, b = x, y

	for b != 0 {
        a, b = b, a % b
    }

    kpk := (x * y) / a

    fmt.Println(kpk)
}
