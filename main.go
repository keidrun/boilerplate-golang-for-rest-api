package main // import "github.com/keidrun/boilerplate-golang-for-rest-api"

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/keidrun/boilerplate-golang-for-rest-api/controllers"
	"github.com/keidrun/boilerplate-golang-for-rest-api/driver"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

const port = 3000

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/users", controller.TokenVerifyMiddleWare(controller.GetUsers())).Methods("GET")

	log.Printf("Server up on port %v....\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
