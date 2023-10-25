package main

import "fmt"

func main() {
    var x, sum int;

	sum = 0
	for i := 1; i <= 5; i++ {
		x = -1
		for x < 0 || x > 200 {
			fmt.Print("Hari ", i, " : ")
			fmt.Scan(&x)
		}
		sum += x
	}

	fmt.Println("Jumlah pengunjung :", sum)
    
}
