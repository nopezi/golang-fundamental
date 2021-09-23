package models

import (
	"fmt"
	belajarrestapi "golang-fundamental/belajar_rest_api"
	"golang-fundamental/belajar_rest_api/structs"
)

func GetPosting(id_kategori string) []structs.Posting {
	dataPosting := []structs.Posting{}
	fmt.Println("cek id kategori", id_kategori)

	db, err := belajarrestapi.KonekMysql()

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