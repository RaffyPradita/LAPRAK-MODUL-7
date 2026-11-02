package main

import "fmt"

func main3() {
	var j, alas, tinggi, n int
	var luas float64
	fmt.Scan(&n)
	for j = 1; j <= n; j += 1 {
		fmt.Scan(&alas, &tinggi)
		luas = 0.5 * float64(alas*tinggi)
		fmt.Println(luas)
	}
}

/*program LuasSegitiga
kamus
	j, alas, tinggi, n : integer
	luas : real
algoritma:
	input (n)
	for j <- 1 to n do:
		input (alas, tinggi)
		luas = 0.5 * alas * tinggi
		output (luas)
	endfor
endprogram*/
