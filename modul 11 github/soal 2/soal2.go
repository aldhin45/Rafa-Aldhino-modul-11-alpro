package main

import (
	"fmt"
)

func main() {
	var suaraMasuk, suaraSah int
	var suara [21]int //231102023
	var calon int

	fmt.Println("Masukkan data suara (0 untuk selesai):")
	for {
		fmt.Scan(&calon)
		if calon == 0 {
			break
		}
		suaraMasuk++
		if calon >= 1 && calon <= 20 {
			suara[calon]++
			suaraSah++
		}
	}

	fmt.Println("\nHasil penghitungan suara:")
	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	var ketua, wakil int
	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua //231102023
			ketua = i
		} else if suara[i] > suara[wakil] && i != ketua {
			wakil = i
		}
	}

	fmt.Println("\nHasil Pemilihan:")
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil Ketua RT:", wakil)
}
