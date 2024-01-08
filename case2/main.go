package main

import "fmt"

type orang struct {
	Nama string
	Umur int
}

type hewan struct {
	Nama string
	Umur int
}

func (p orang) berbicara() {
	fmt.Printf("Halo Nama saya %s, saya berumur %d. \n", p.Nama, p.Umur)
}

func (h hewan) berbicara() {
	fmt.Printf("Halo Nama saya %s, berumur %d. \n", h.Nama, h.Umur)
}

type MakhlukHidup interface {
	berbicara()
}

func speak(tipe MakhlukHidup) {
	fmt.Println("Saya Makhluk Hidup")
}

func main() {
	m1 := orang{
		Nama: "Fajrin",
		Umur: 25,
	}
	m2 := hewan{
		Nama: "kirby",
		Umur: 10,
	}
	m1.berbicara()
	m2.berbicara()
	speak(m1)
	speak(m2)
}
