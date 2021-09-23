package structs

type Users struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
}

type Result struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Posting struct {
	Id int `json:"id"`
	Id_kategori int `json:"id_kategori"`
	Judul string `json:"judul"`
	Isi string `json:"isi"`
	gambar string `json:"gambar"`
	Created_at string `json:"created_at"`
}