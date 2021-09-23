package models

import (
	"fmt"

	belajarrestapi "golang-fundamental/belajar_rest_api"
	"golang-fundamental/belajar_rest_api/structs"
)

func CobaModel()  {
	fmt.Println("masok")
}

func GetUser() []structs.Users {
	// var user Users
	// var users structs.Users
	users := []structs.Users{}
	db, err := belajarrestapi.KonekMysql()

	if err != nil {
		fmt.Println("error gan", err)
	}
	db.Find(&users)
	return users
}