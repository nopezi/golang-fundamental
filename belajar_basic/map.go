package belajarbasic

import "fmt"

func MapPertama() {

	var myMap map[string]int
	myMap = map[string]int{}

	myMap["makanan"] = 205

	fmt.Println("map pertama ", myMap["makanan"])
}

func MapKedua() {

	var myMap map[string]string
	myMap = map[string]string{"makanan": "rendang", "minuman": "air zam zam"}

	fmt.Println("map kedua ", myMap["minuman"])

}

func MapKetiga() {
	myMap := map[string]string{
		"makanan": "rendang",
		"minuman": "juz strowberry",
	}

	fmt.Println("Map ketiga ", myMap)
}

func MapKeempat() {
	myMap := map[string]string{
		"makanan": "nasi liwet",
		"minuman": "es teh",
		"buah":    "jambu",
	}

	for key, value := range myMap {
		fmt.Println("key : ", key, " value : ", value)
	}

	// untuk hapus data array
	delete(myMap, "buah")
	fmt.Println("setelah di hapus ")

	for key, value := range myMap {
		fmt.Println("key : ", key, " value : ", value)
	}
}

func SliceMapPertama() {
	myMap := []map[string]string{
		{"makanan": "gudeg", "minuman": "es teh"},
		{"makanan": "mie ayam", "minuman": "es jeruk"},
	}

	for _, v := range myMap {
		fmt.Println("makanan ", v["makanan"], " - Minuman ", v["minuman"])
	}
}

func MapQuiz() {
	// hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int

	rataRata := 0
	for _, v := range scores {
		rataRata += v
		if v >= 90 {
			goodScores = append(goodScores, v)
		}
	}

	length := float64(len(scores))
	average := float64(rataRata)
	hasil := average / length

	fmt.Println("nilai rata rata map ", hasil)
	fmt.Println("skor lebih dari 90 : ", goodScores)

}
