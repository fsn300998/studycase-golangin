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
	data = append(data, m1)
	for _, v := range data {
		fmt.Printf("Halo nama saya %s, umur saya %d tahun, dan alamat saya di %s", v.Nama, v.Umur, v.Alamat)
	}
}
