package belajarbasic

import "fmt"

func SlicePertama() {

	var makanan []string

	makanan = append(makanan, "rendang")
	makanan = append(makanan, "sate")
	makanan = append(makanan, "bakso")

	fmt.Println("slice pertama", makanan)

	for _, jenis := range makanan {
		fmt.Println("slice pertama ake for ", jenis)
	}

}

func SliceKedua() {

	makanan := []string{"rendang", "bakso", "sate"}

	fmt.Println("Slice kedua ", makanan)

}
