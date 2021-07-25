package belajaroracle

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	// goracle "gopkg.in/goracle.v2"
)

func Konek() {
	// db, err := sql.Open("godror", `user="dev" password="dev" connectString="36.91.39.115:1521/orclpdb1"`)
	const dsn = `user="rsnew"
				 password="rsnew"
				 connectString="36.91.39.115:1521/xe"`
	db, err := sql.Open("godror", dsn)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	println("oke konek gays")

	err = db.Ping()
	// os.Exit(1)
	if err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}

	// query := `select id, username from pengguna`
	query := `select * from kunjungan`

	// query := `select sysdate from dual`
	fmt.Println("query nya ", query)

	rows, err := db.Query(query)

	// rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var nomor, nama string

	for rows.Next() {
		rows.Scan(&nomor, &nama)
		fmt.Println("nomor ", nomor, " nama ", nama)
	}
	os.Exit(1)
	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}

func Konek2() {

	username := "dev"
	password := "dev"
	host := "36.91.39.115"
	database := "xe"

	db, err := sql.Open("goracle", username+"/"+password+"@"+host+"/"+database)

	// os.Exit(1)
	if err != nil {
		fmt.Println("... DB Setup Failed")
		// fmt.Println(err)
		return
	}
	defer db.Close()

}

func Konek3() {

	username := "dev"
	password := "dev"
	db, err := sqlx.Connect("godror", username+"/"+password+"@36.91.39.115:1521/XE")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("berhasil konek")
	tx := db.MustBegin()

	query := `select * from pengguna`
	tx.MustExec(query)
}
