package main

import (
    "fmt"
)

func main() {
    var h1, m1, h2, m2, h, m int;

    fmt.Scan(&h1, &m1, &h2, &m2)
	fmt.Println("");

	if h1 <= h2 && m1 <= m2 {
		h = h2 - h1;
		m = m2 - m1;
	} else if m1 <= m2{
		h = h2 - h1 + 12;
		m = m2 - m1;
	} else {
		h = h2 - h1 + 12;
		m = m2 - m1 + 60;
	}

	fmt.Println(h,"jam", m, "menit")
}
