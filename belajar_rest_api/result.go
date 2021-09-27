package belajarrestapi

import (
	"encoding/json"
	"golang-fundamental/belajar_rest_api/structs"
	"net/http"
)

func ResultOk(w http.ResponseWriter, res structs.Result) {
	hasil, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(hasil)
}
