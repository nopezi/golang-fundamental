package controllers

import (
	"encoding/json"
	"fmt"
	"golang-fundamental/belajar_rest_api/models"
	"golang-fundamental/belajar_rest_api/structs"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "berhasil")
}

func UserController(w http.ResponseWriter, r *http.Request) {
	var res structs.Result
	data := models.GetUser()
	res.Status = "true"
	res.Message = "Berhasil"
	res.Data = data
	// res := structs.Result(Status: "true", Message: "Berhasil", Data: data)
	hasil, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(hasil)
}