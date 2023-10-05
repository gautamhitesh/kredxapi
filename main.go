package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/gautamhitesh/kredxapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World. Starting golang server")

	r := mux.NewRouter().StrictSlash(true)
	// r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/signup", SignUp).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")
	r.HandleFunc("/sessions", SessionLogs).Methods("GET")
	r.HandleFunc("/borrow", Borrow).Methods("POST")
	r.HandleFunc("/lend", Lend).Methods("POST")
	r.HandleFunc("/report", Report).Methods("GET")
	// log.Fatal(http.ListenAndServe(":3200", r))
	http.Handle("/", r)
	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())

}
