package controllers

import (
	"encoding/json"
	"fmt"
	belajarrestapi "golang-fundamental/belajar_rest_api"

	// "golang-fundamental/belajar_rest_api/configs"
	"golang-fundamental/belajar_rest_api/models"
	"golang-fundamental/belajar_rest_api/structs"
	"io/ioutil"
	"net/http"
)

func Posting(w http.ResponseWriter, r *http.Request) {
	if !belajarrestapi.Auth(w, r) {
		return
	}
	var res structs.Result

	idkategori := r.FormValue("id_kategori")

	data := models.GetPosting(idkategori)

	res.Status = "true"
	res.Message = "berhasil mendapatkan data posting"
	res.Data = data

	belajarrestapi.ResultOk(w, res)
	// configs.ResultOk(w, res)
}

func TambahPosting(w http.ResponseWriter, r *http.Request) {
	if !belajarrestapi.Auth(w, r) {
		return
	}

	var res structs.Result
	var posting structs.TambahPosting
	input, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(input, &posting)

	res.Data = nil
	res.Status = "false"

	if posting.Id_kategori == 0 {
		res.Message = "id kategori tidak boleh kosong"
	} else if posting.Judul == "" {
		res.Message = "judul tidak boleh kosong"
	} else if posting.Isi == "" {
		res.Message = "isi tidak boleh kosong"
	} else {
		res.Status = "true"
		res.Message = "berhasil tambah data"
		res.Data = posting
	}

	models.TambahPosting(posting)

	belajarrestapi.ResultOk(w, res)
}

func UpdatePosting(w http.ResponseWriter, r *http.Request) {
	if !belajarrestapi.Auth(w, r) {
		return
	}
	var res structs.Result

	input, _ := ioutil.ReadAll(r.Body)
	var inputData structs.UpdatePosting
	json.Unmarshal(input, &inputData)
	hasilUpdate := models.UpdatePosting(inputData)
	fmt.Println("vars update ", inputData)

	res.Status = "true"
	res.Message = "berhasil update data"
	res.Data = hasilUpdate

	dataUpdate, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataUpdate)
}
