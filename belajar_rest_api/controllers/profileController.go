package controllers

import (
	"encoding/json"
	"fmt"
	belajarrestapi "golang-fundamental/belajar_rest_api"
	"io/ioutil"

	// "golang-fundamental/belajar_rest_api/configs"
	"golang-fundamental/belajar_rest_api/models"
	"golang-fundamental/belajar_rest_api/structs"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request)  {
	var res structs.Result

	data := models.GetProfile()

	res.Status = "true"
	res.Message = "Berhasil mendapatkan data"
	res.Data = data
	belajarrestapi.ResultOk(w, res)
	// configs.ResultOk(w, res)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request)  {
	var res structs.Result
	var updateProfile structs.UpdateProfile

	input, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(input, &updateProfile)
	hasil := models.UpdateProfile(updateProfile)

	fmt.Println("update profile input ", updateProfile)
	res.Status = "true"
	res.Message = "berhasil update data profile"
	res.Data = hasil

	belajarrestapi.ResultOk(w, res)
}