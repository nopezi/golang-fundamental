package belajarbasic

import "fmt"

type Kostumer struct {
	Nama, Alamat string
	Umur int
}

func BelajarStruct()  {
	var nopezi Kostumer
	nopezi.Nama = "Nopezi"
	nopezi.Alamat = "Yogyakarta"
	nopezi.Umur = 27

	fmt.Println("hasil struct ", nopezi)

	// versi kedua
	jono := Kostumer{
		Nama: "Jono",
		Alamat: "bengkulu",
		Umur: 31,
	}

	fmt.Println("hasil struct 2 ", jono)

	// versi ketiga
	bambang := Kostumer{"bambang", "jakarta", 13}

	fmt.Println("hasil struct 3 ", bambang)
}

func (customer Kostumer) sayHalo(nama string) {
	fmt.Println("hai", nama, "nama saya", customer.Nama)
}

func BelajarStructMethod()  {
	var nopezi Kostumer
	nopezi.Nama = "Nopezi"
	nopezi.Alamat = "Yogyakarta"
	nopezi.Umur = 27

	nopezi.sayHalo("jono")
}