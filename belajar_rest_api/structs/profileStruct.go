package structs

type Profile struct {
	Id int `json:"id"`
	Nama_lengkap string `json:"nama_lengkap"`
	Tempat_lahir string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Agama string `json:"agama"`
	Email string `json:"email"`
	No_hp string `json:"no_hp"`
	Hobi string `json:"hobi"`
	Deskripsi string `json:"deskripsi"`
	Url_gambar string `json:"url_gambar"`
}

type UpdateProfile struct {
	Id int `json:"id"`
	Nama_lengkap string `json:"nama_lengkap"`
	// Tempat_lahir string `json:"tempat_lahir"`
	// Tanggal_lahir string `json:"tanggal_lahir"`
	// Agama string `json:"agama"`
	// Email string `json:"email"`
	Deskripsi string `json:"deskripsi"`
}