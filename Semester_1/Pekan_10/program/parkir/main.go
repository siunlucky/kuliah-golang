package main

import (
    "fmt"
)

func main() {
    var h1, m1, h2, m2, h, m int;

    fmt.Scan(&h1, &m1, &h2, &m2)
	fmt.Println("");

	if h1 == 6 || h2 == 6 || (h2 == 5 && m2 > 0) || (h1 == 5 && m1 > 0) {
		fmt.Println("Parkiran sedang tutup")
	} else {
		if h1 <= h2 { 
			if m1 > m2 {
				h2 -= 1
				m2 += 60
			}
			h = h2 - h1;
			m = m2 - m1;
		} else {
			if m1 > m2 {
				h2 -= 1
				m2 += 60
			}
			h = h2 - h1 + 12;
			m = m2 - m1;
		}
		fmt.Println(h,"jam", m, "menit")
	}
}
