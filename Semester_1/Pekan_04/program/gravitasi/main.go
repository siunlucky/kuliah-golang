package main

import "fmt";

func main() {
	var merkurius, venus, mars, bumi, yupiter, saturnus, uranus, neptunus, x float64
	bumi = 9.8

	merkurius = 0.38 * bumi
	venus = 0.91 * bumi
	mars = 0.38 * bumi
	yupiter = 2.37 * bumi
	saturnus = 0.92 * bumi
	uranus = 0.89 * bumi
	neptunus = 1.13 * bumi

	fmt.Scan(&x)

	fmt.Printf("%.0f %.0f %.0f %.0f %.0f %.0f %.0f %.0f\n", x * merkurius, x * venus, x * bumi, x * mars, x * yupiter, x * saturnus, x * uranus, x * neptunus)

}