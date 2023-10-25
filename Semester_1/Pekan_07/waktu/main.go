package main

import (
	"fmt"
);

func main()  {
	var s, j, m, d int;

	fmt.Scan(&s);
	fmt.Println("");

	j =  s / 3600
	s = s % 3600

	m = s / 60

	d = s % 60

	fmt.Println(j, "jam", m, "menit dan", d, "detik");


}