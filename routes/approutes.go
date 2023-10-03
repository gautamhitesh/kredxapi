package routes

import (
	"encoding/json"
	"io"
	"net/http"

	. "github.com/gautamhitesh/kredxapi/models"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World, server is running on port 3200")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// username := vars["username"]
	reqBody, _ := io.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	json.NewEncoder(w).Encode(user.Username + " Signed Up")
}

func Login(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	json.NewEncoder(w).Encode(user.Username + " Logged in!")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Logged out!")
}

func Lend(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Lent!")
}
func Borrow(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Borrowed!")
}

func Report(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Report!")
}
