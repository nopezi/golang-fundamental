package calculation

// nama function huruf besar bisa di ambil beda package
func Add(number int, numberTwo int) int {
	return add(number, numberTwo)
}

// nama function awal huruf kecil hanya dapat di ambil di satu package
func add(number int, numberTwo int) int {
	return number + numberTwo
}
