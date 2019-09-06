package main

import "strings"

//Calc Adalah
type Calc struct {
	BilanganSatu int
	BilanganDua  int
	operasi      string
}

func (c *Calc) setSatu(angka int) {
	c.BilanganSatu = angka
}
func (c *Calc) setDua(angka int) {
	c.BilanganDua = angka
}
func (c *Calc) setOps(str string) {
	c.operasi = str
}
func (c Calc) hitung() (output int) {
	// var hasil int
	if c.operasi == "+" {
		output = c.BilanganSatu + c.BilanganDua
		return
	} else if c.operasi == "-" {
		output = c.BilanganSatu - c.BilanganDua
		return
	} else if c.operasi == "*" || strings.ToLower(c.operasi) == "x" {
		output = c.BilanganSatu * c.BilanganDua
		return
	} else {
		output = c.BilanganSatu / c.BilanganDua
		return
	}
}
