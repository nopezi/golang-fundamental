package belajarbasic

func If(umur int) string {
	var hasil string

	if umur < 26 {
		hasil = "masih mahasiswa"
	} else if umur > 26 {
		hasil = "udah perjaka tua"
	} else {
		hasil = "pas sekarang"
	}

	return hasil

}

// if
// if else
// else if
// if bersarang
