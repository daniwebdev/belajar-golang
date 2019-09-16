package main

import "strings"

//Calc Adalah
type Calc struct {
	BilanganSatu int64
	BilanganDua  int64
	operasi      string
}

func (c *Calc) setSatu(angka int64) {
	c.BilanganSatu = angka
}
func (c *Calc) setDua(angka int64) {
	c.BilanganDua = angka
}
func (c *Calc) setOps(str string) {
	c.operasi = str
}
func (c Calc) hitung() (output int64) {
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
