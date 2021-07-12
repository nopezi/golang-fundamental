package belajarbasic

func SwitchCase(anggota string) string {
	var hasil string
	switch anggota {
	case "ayah":
		hasil = "abdul muluk"
	case "ibu":
		hasil = "awaliyah"
	case "anak1":
		hasil = "nopezi"
	case "anak2":
		hasil = "pera"
	case "anak3":
		hasil = "nadia"
	default:
		hasil = "orang lain"
	}
	return hasil
}
