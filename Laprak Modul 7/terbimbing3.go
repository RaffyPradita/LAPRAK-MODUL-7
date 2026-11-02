package main

import "fmt"

func main() {
	var dinar, dirham, fals, qirat int
	fmt.Scan(&qirat)
	dinar = qirat / 600
	dirham = qirat % 600 / 60
	fals = qirat % 600 % 60 / 6
	qirat = qirat % 600 % 60 % 6
	fmt.Println(dinar, dirham, fals, qirat)
}

/*program MataUang
kamus
	dinar, dirham, fals, qirat : integer
algoritma
	input (qirat)
	dinar <- qirat / 600
	dirham <- qirat % 600 / 60
	fals <- qirat % 600 % 60 / 6
	qirat <- qirat % 600 % 60 % 6
	output (dinar, dirham, fals, qirat)
endprogram*/
