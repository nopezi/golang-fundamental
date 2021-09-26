package structs

type Posting struct {
	Id int `json:"id"`
	Id_kategori int `json:"id_kategori"`
	Judul string `json:"judul"`
	Isi string `json:"isi"`
	// gambar string `json:"gambar"`
	Created_at string `json:"created_at"`
}

type TambahPosting struct {
	Id_kategori int `json:"id_kategori"`
	Judul string `json:"judul"`
	Isi string `json:"isi"`
	// gambar string `json:"gambar"`
}

type UpdatePosting struct {
	Id int `json:"id"`
	Id_kategori int `json:"id_kategori"`
	Judul string `json:"judul"`
	Isi string `json:"isi"`
}