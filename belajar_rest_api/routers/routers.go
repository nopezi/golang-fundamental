package routers

import (
	"golang-fundamental/belajar_rest_api/controllers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Routers() {
	log.Println("server berjalan di http://127.0.0.1:12345")
	
	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/", controllers.HomeController)
	route.HandleFunc("/users", controllers.UserController).Methods("get")
	route.HandleFunc("/posting", controllers.Posting).Methods("get")

	log.Fatal(http.ListenAndServe(":12345", route))

}
