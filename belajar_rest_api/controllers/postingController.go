package controllers

import (
	"encoding/json"
	belajarrestapi "golang-fundamental/belajar_rest_api"
	"golang-fundamental/belajar_rest_api/models"
	"golang-fundamental/belajar_rest_api/structs"
	"net/http"
)

func Posting(w http.ResponseWriter, r *http.Request)  {
	if !belajarrestapi.Auth(w, r){ return }
	var res structs.Result

	idkategori := r.FormValue("id_kategori")

	data := models.GetPosting(idkategori)

	res.Status = "true"
	res.Message = "berhasil mendapatkan data posting"
	res.Data = data

	hasil, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(hasil)
}