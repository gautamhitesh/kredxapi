package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gautamhitesh/kredxapi/middleware"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World, server is running on port 3200")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	middleware.SignUp(w, r)
}

func Login(w http.ResponseWriter, r *http.Request) {
	middleware.LoginHandler(w, r)

	// json.NewEncoder(w).Encode(user.Username + " Logged in!")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	middleware.LogoutHandler(w, r)
	// json.NewEncoder(w).Encode("Logged out!")
}

func Lend(w http.ResponseWriter, r *http.Request) {
	if middleware.BasicAuth(r) {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode("Lent")
	} else {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Please sign in to continue")
	}
}
func Borrow(w http.ResponseWriter, r *http.Request) {
	if middleware.BasicAuth(r) {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode("Borrowed")
	} else {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Please sign in to continue")
	}
}

func Report(w http.ResponseWriter, r *http.Request) {
	if middleware.BasicAuth(r) {
		json.NewEncoder(w).Encode("Report")
	}
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode("Please sign in to continue")
}

func SessionLogs(w http.ResponseWriter, r *http.Request) {
	middleware.SessionLogs(w, r)
}
