package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/gautamhitesh/kredxapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World. Starting golang server")

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/signup", SignUp).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/login", Logout).Methods("POST")
	r.HandleFunc("/borrow", Borrow).Methods("POST")
	r.HandleFunc("/lend", Lend).Methods("POST")
	r.HandleFunc("/report", Report).Methods("GET")
	log.Fatal(http.ListenAndServe(":3200", r))
}
