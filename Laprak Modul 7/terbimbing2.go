package main

import "fmt"

func main2() {
	var N, j int
	fmt.Scan(&N)
	for j = 1; j <= N; j += 1 {
		fmt.Print(j*j, " ")
	}
}

/*program Kuadrat
kamus
	N, j : integer
algoritma
	input (N)
	for j = 1 to N do
		output (j*j, " ")
	endfor
endprogram*/
