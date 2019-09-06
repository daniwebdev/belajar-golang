package main

//Hero is
type Hero struct {
	name  string
	power int
}

func (h *Hero) attact() {
	h.power--
}
