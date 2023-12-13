package main

import "fmt";

func main() {
	var x, y int
	var isRoot bool

	fmt.Scan(&x, &y)

	isRoot = (x * x * x) == y

	fmt.Println(isRoot);
}