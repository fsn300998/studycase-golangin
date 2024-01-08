package main

import "fmt"

type identitas struct {
	Nama   string
	Umur   int
	Alamat string
}

var data []identitas

func main() {
	m1 := identitas{
		Nama:   "Fajrin",
		Umur:   25,
		Alamat: "Tasikmalaya",
	}
	m2 := identitas{
		Nama:   "Susanto",
		Umur:   24,
		Alamat: "Solo",
	}
	data = append(data, m1, m2)
	for _, v := range data {
		fmt.Printf("Halo nama saya %s, umur saya %d tahun, dan alamat saya di %s.\n", v.Nama, v.Umur, v.Alamat)
	}
}
