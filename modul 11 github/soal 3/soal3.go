package main

import (
	"fmt"
	"math/rand"
)

const NMAX = 1000000

var data [NMAX]int //231102023

func isiArray(n int) {
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(1000)
	}
}

func posisi(n, k int) int {
	for i := 0; i < n; i++ {
		idx := rand.Intn(n)
		if data[idx] == k {
			return idx
		}
	}
	return -1
}

func main() {
	var n, k int
	fmt.Scanln(&n, &k)

	isiArray(n)
	fmt.Print("Data: ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println() //231102023

	idx := posisi(n, k)
	if idx != -1 {
		fmt.Println("Angka", k, "berada pada posisi ke", idx+1)
	} else {
		fmt.Println("TIDAK ADA")
	}
}
