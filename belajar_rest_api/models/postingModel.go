package models

import (
	"fmt"
	belajarrestapi "golang-fundamental/belajar_rest_api"
	// "golang-fundamental/belajar_rest_api/configs"
	"golang-fundamental/belajar_rest_api/structs"
)

func GetPosting(id_kategori string) []structs.Posting {
	dataPosting := []structs.Posting{}
	// fmt.Println("cek id kategori", id_kategori)

	db, err := belajarrestapi.KonekMysql()
	// db, err := configs.KonekMysql()

	if err != nil {
		fmt.Println("error get posting", err)
	}

	if id_kategori == "" {
		sql := `SELECT
				id, id_kategori, judul, isi, gambar, created_at
			FROM
				posting
			ORDER BY
				id
			DESC`

		db.Raw(sql).Scan(&dataPosting)
	} else {
		sql := `SELECT
				id, id_kategori, judul, isi, gambar, created_at
			FROM
				posting
			WHERE
				id_kategori = ` + id_kategori +
			` ORDER BY
				id
			DESC`

		db.Raw(sql).Scan(&dataPosting)
	}

	return dataPosting
}

func TambahPosting(dposting structs.TambahPosting) {
	fmt.Println("data tambah posting", dposting)

	db, _ := belajarrestapi.KonekMysql()
	// db, _ := configs.KonekMysql()
	db.Table("posting").Create(&dposting)
}

func UpdatePosting(input structs.UpdatePosting) structs.Posting {
	fmt.Println("data update", input)
	var seluruh structs.Posting
	// dataUpdate := map[string]interface{}{
	// 	"id_kategori": input.Id_kategori,
	// }

	db, _ := belajarrestapi.KonekMysql()
	// db, _ := configs.KonekMysql()
	db.Table("posting").Updates(&input)
	db.Table("posting").Where("id = ?", input.Id).Find(&seluruh)
	return seluruh
}
