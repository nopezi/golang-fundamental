package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB
var err error

const USERNAME = "admin"
const PASSWORD = "admin"

func Auth(w http.ResponseWriter, r *http.Request) bool {
    username, password, ok := r.BasicAuth()
    if !ok {
        w.Write([]byte(`something went wrong`))
        return false
    }

    isValid := (username == USERNAME) && (password == PASSWORD)
    if !isValid {
        // w.Write([]byte(`wrong username/password`))
		fmt.Fprintln(w, "basic auth tidak boleh kosong")
        return false
    }

    return true
}

func CekAuth(w http.ResponseWriter, r *http.Request)  {
	if !Auth(w, r)         { return }
}

func ConfigMysql() (*gorm.DB, error) {
	dbname := "heroku_13b21ebb9b93855"
	host := "us-cdbr-east-02.cleardb.com"
	user := "b10250d4524049"
	pass := "d377e116"
	port := "3306"
	konek := user + `:` + pass + `@tcp(`+ host + `:` + port +`)/` + dbname +`?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := gorm.Open(mysql.Open(konek), &gorm.Config{})

	if err != nil {
		fmt.Println("db error", err.Error())
	}

	return db, err
}

func KonekMysql() (*gorm.DB, error) {
	db, err := ConfigMysql()
	return db, err
}