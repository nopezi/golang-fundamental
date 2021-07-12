package belajarbasic

import "fmt"

func ArrayPertama() {

	var bahasa [5]string
	bahasa[0] = "PHP"
	bahasa[1] = "JAVASCRIPT"
	bahasa[2] = "PYTHON"
	bahasa[3] = "GOLANG"

	fmt.Println(bahasa)
	panjang := len(bahasa)
	fmt.Println("total array ", panjang)

}

func ArrayKedua() {

	bahasa := [5]string{"php", "javascript", "python", "golang"}
	fmt.Println("array versi kedua ", bahasa)

}

func ArrayKetiga() {

	bahasa := [...]string{
		"php",
		"javascript",
		"python",
		"golang",
		"C#",
	}
	fmt.Println("array versi ketiga ")
	for index, v := range bahasa {
		fmt.Println("index ; ", index, " bahasa : ", v)
	}

}
