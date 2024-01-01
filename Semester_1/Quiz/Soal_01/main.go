package main

import "fmt"

func main() {
	var isRainy, haveUmbrella string

	fmt.Scan(&isRainy, &haveUmbrella)

	if isRainy == "tidak" || (isRainy == "ya" && haveUmbrella == "ya") {
		fmt.Println("berangkat")
	} else {
		fmt.Println("diam di rumah")
	}
}
