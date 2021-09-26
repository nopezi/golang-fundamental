package belajarbasic

import "fmt"

// defer untuk tetap menjalankan function walaupun muncul error
// panic untuk menampilkan error dan menghentikan program
// recover giunakan untuk menangkap data panic supaya program tidak berhenti

func logging()  {
	pesanRecover := recover()
	fmt.Println("pesan recover ", pesanRecover)
	fmt.Print("selesai memanggil function")
}

func DeferPanicRecover(error bool)  {
	defer logging()
	if error {
		fmt.Println("")
		panic("Aplikasi error panik ")
	}
	fmt.Println("Aplikasi jalan")
}