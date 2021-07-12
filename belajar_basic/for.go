package belajarbasic

import "fmt"

func ForPertama(total int) {
	for i := 0; i < total; i++ {
		fmt.Println("belajar go perulangan ", i)
	}
}

func ForKedua(total int) {
	i := 1
	for i <= total {
		fmt.Println("kedua ", i)
		i++
	}
}

func ForKetiga(kalimat string) {
	for _, letter := range kalimat {
		fmt.Println("kalimat : ", letter)
	}
}

func KuisFor(kalimat string) {
	title := "Golang the best language"
	for cek, kata := range title {
		if cek%2 == 0 || cek == 0 {
			fmt.Println("kuis for hanya print index ", cek, " huruf yang genap ", string(kata))
		}
	}
}

func KuisForDua() {
	title := "Golang the best language"
	for _, v := range title {
		switch string(v) {
		case "a", "i", "u", "o":
			fmt.Println(" huruf ", string(v), " index ke ", v)
		}
	}
}
