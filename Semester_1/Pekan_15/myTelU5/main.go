package main

import "fmt"

func main() {
    var x, sum, sumGold, sumPlatinum, sumSilver, countGold, countPlatinum, countSilver int;
	var averageGold, averagePlatinum, averageSilver float64;

	sum, sumGold, sumPlatinum, sumSilver = 0, 0, 0, 0
	
	for sum < 500 {
		x = 0
		for x < 50 {
			fmt.Scan(&x)
		}

		sum += x

		if x > 200 {
			countGold++
			sumGold += x
		} else if x >= 100 && x <= 200{
			countPlatinum++
			sumPlatinum += x
		} else if x >= 50 && x < 100{
			countSilver++
			sumSilver += x
		}
	}

	averageGold = float64(sumGold) / float64(countGold)
	averagePlatinum = float64(sumPlatinum) / float64(countPlatinum)
	averageSilver = float64(sumSilver) / float64(countSilver)

	if averageGold != averageGold {
		averageGold = 0

		fmt.Print("Gold user : ", averageGold)
	} else {
		fmt.Printf("Gold user : %.2f", averageGold)
	}

	if averagePlatinum != averagePlatinum {
		averagePlatinum = 0

		fmt.Print(", Platinum user : ", averagePlatinum)
	} else {
	
		fmt.Printf(", Platinum user : %.2f", averagePlatinum)
	}

	if averageSilver != averageSilver {
		averageSilver = 0

		fmt.Print(", Silver user : ", averageSilver)
	} else {
		fmt.Printf(", Silver user : %.2f", averageSilver)
	}

}
