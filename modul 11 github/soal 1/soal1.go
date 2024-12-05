package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var suaraMasuk, suaraSah, calon int
	var suara [21]int //231102023

	fmt.Print("Masukkan data suara: ")
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

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	for i := 0; i < 20; i++ { //231102023
		calonAcak := rand.Intn(21)
		if calonAcak >= 1 && calonAcak <= 20 && suara[calonAcak] > 0 {
			fmt.Println(calonAcak, ":", suara[calonAcak])
			suara[calonAcak] = 0
		} else {
			i--
		}
	}
}
