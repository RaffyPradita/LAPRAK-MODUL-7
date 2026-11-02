package main

import "fmt"

func main1() {
	var x, y, j, jumlah int
	fmt.Scan(&x, &y)
	jumlah = 0
	for j = x; j <= y; j += 1 {
		jumlah = jumlah + j
	}
	fmt.Println(jumlah)
}

/*program JumlahBilangan
kamus
	x, y, j, jumlah : integer

algoritma
	input (x, y)
	jumlah <- 0
	for j <- x to y do
		output (j)
		jumlah <- jumlah + j
	endfor
	output(jumlah)
endprogram*/
