package configs

import (
	"golang-fundamental/belajar_rest_api/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Routers() {
	// log.Println("server berjalan di http://127.0.0.1:80")
	
	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/", controllers.HomeController)
	route.HandleFunc("/users", controllers.UserController).Methods("get")

	route.HandleFunc("/posting", controllers.Posting).Methods("get")
	route.HandleFunc("/tambah_posting", controllers.TambahPosting).Methods("post")
	route.HandleFunc("/update_posting", controllers.UpdatePosting).Methods("put")

	route.HandleFunc("/profile", controllers.GetProfile).Methods("get")
	route.HandleFunc("/profile/update", controllers.UpdateProfile).Methods("put")

	// log.Fatal(http.ListenAndServe("0.0.0.0:", nil))
	// log.Fatal(http.ListenAndServe("0.0.0.0:8080", route))
	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}
	jalan := http.ListenAndServe("0.0.0.0:"+port, route)
	log.Println("server berjalan di ", jalan)
	log.Fatal(jalan)

}
