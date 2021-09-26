package models

import (
	belajarrestapi "golang-fundamental/belajar_rest_api"
	// configsdb "golang-fundamental/belajar_rest_api/configs"
	"golang-fundamental/belajar_rest_api/structs"
)

func GetProfile() structs.Profile {
	var profile structs.Profile
	db, _ := belajarrestapi.KonekMysql()
	// db, _ := configs.KonekMysql()

	sql := `SELECT 
				id, nama_lengkap, tempat_lahir, tanggal_lahir,
				agama, email, deskripsi, 
			CONCAT
				('http://shielded-beyond-23529.herokuapp.com/gambar/', foto) as url_gambar
			FROM
				profil`
	db.Raw(sql).Scan(&profile)

	return profile
}

// update data profiile
// input = data yang di ambil dari request body
// resturn data berupa struct

func UpdateProfile(input structs.UpdateProfile) structs.Profile {
	db, _ := belajarrestapi.KonekMysql()
	db.Table("profil").Updates(&input)
	hasil := GetProfile()
	return hasil
}