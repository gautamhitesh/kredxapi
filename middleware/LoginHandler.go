package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gautamhitesh/kredxapi/models"
	"github.com/google/uuid"
)

type User struct {
	models.User
}

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

var users = map[string]string{
	"admin":   "password",
	"mac":     "air",
	"android": "dellzte",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Loggging in")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Data allowed is only in the form of URL endcoded form", http.StatusBadRequest)
		return
	}
	password := r.PostForm.Get("password")
	username := r.PostForm.Get("username")

	fmt.Println(users)

	if pass, ok := users[username]; ok {
		// session, _ := store.Get(r, "session.id")
		if pass == password {
			// session.Values["authenticated"] = true
			// session.Save(r, w)
			sessionToken := uuid.NewString()
			expiresAt := time.Now().Add(600 * time.Second)
			Sessions[sessionToken] = Session{
				username: username,
				expiry:   expiresAt,
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: expiresAt,
			})
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Login Successfully"))
}

func SessionLogs(w http.ResponseWriter, r *http.Request) {
	// session, _ := store.Get(r, "session.id")
	// if session.Values["authenticated"] != nil {
	// 	w.Write([]byte(time.Now().String()))
	// } else {
	// 	http.Error(w, "Forbidden", http.StatusForbidden)
	// }
	sessions := []string{}
	for _, v := range Sessions {
		fmt.Println(v)
		sessions = append(sessions, fmt.Sprintf("%#v", v))
	}
	json.NewEncoder(w).Encode(sessions)
}
