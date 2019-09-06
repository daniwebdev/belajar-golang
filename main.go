package main

import "fmt"

func main() {

	hit := Calc{1, 2, "+"}
	hit.setSatu(5)
	hit.setDua(10)

	fmt.Printf("Hasil nya : %d", hit.hitung())
}
