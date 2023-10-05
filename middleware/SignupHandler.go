package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gautamhitesh/kredxapi/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Wrong body")
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	username := user.Username
	password := user.Password

	_, exists := users[username]
	if exists {
		w.Write([]byte("User already exists, please login!"))
		return
	}
	users[username] = password
	fmt.Println("Users")
	for _, i := range users {
		fmt.Println(i)
	}

	w.Write([]byte("Signed up successfuly"))

}
