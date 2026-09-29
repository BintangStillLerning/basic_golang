package main

import (
	"fmt"
	"strconv"
)

func potongan(uang float64) float64 {
	if uang > 10_000_000 {
		return uang * 0.70
	} else if uang > 1_000_000 {
		return uang * 0.70
	} else if uang > 0 {
		return uang * 0.60
	} else {
		return 0
	}
}

func formatRupiah(angka float64) string {
	n := int(angka)
	s := strconv.Itoa(n)
	panjang := len(s)

	hasil := ""
	for i, c := range s {
		if i != 0 && (panjang-i)%3 == 0 {
			hasil += "."
		}
		hasil += string(c)
	}
	return "Rp " + hasil
}

func main() {
	var uang float64

	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scanln(&uang)

	hasil := potongan(uang)

	fmt.Println("Hasil Bersih:", formatRupiah(hasil))
}
