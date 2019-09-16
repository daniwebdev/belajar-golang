package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	hit := Calc{1, 2, "+"}
	fmt.Print("Nilai 1 => ")
	satu, _ := reader.ReadString('\n')
	satu = strings.Replace(satu, "\n", "", -1)
	set_satu, err := strconv.ParseInt(satu, 10, 64)
	if err != nil {

	}
	hit.setSatu(set_satu)

	fmt.Print("Nilai 2 => ")
	dua, _ := reader.ReadString('\n')
	dua = strings.Replace(dua, "\n", "", -1)
	set_dua, e := strconv.ParseInt(dua, 10, 64)
	if e != nil {

	}
	hit.setDua(set_dua)

	// interface{}
	fmt.Printf("Hasil nya : %d", hit.hitung())
	fmt.Println()
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Simple Shell")
// 	fmt.Println("---------------------")

// 	// for {
// 		fmt.Print("-> ")
// 		text, _ := reader.ReadString('\n')
// 		// convert CRLF to LF
// 		text = strings.Replace(text, "\n", "", -1)

// 		if strings.Compare("hi", text) == 0 {
// 			fmt.Println("hello, Yourself")
// 		}
// 	// }
// }
