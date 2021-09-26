package structs

type Profile struct {
	Id int `json:"id"`
	Nama_lengkap string `json:"nama_lengkap"`
	Tempat_lahir string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Agama string `json:"agama"`
	Email string `json:"email"`
	Deskripsi string `json:"deskripsi"`
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