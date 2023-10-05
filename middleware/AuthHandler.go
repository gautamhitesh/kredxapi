package middleware

import (
	"net/http"
)

func BasicAuth(r *http.Request) bool {
	c, err := r.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			// w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		// w.WriteHeader(http.StatusBadRequest)
		return false
	}
	sessionToken := c.Value

	userSession, exists := Sessions[sessionToken]

	if !exists {
		// w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	if userSession.isExpired() {
		delete(Sessions, sessionToken)
		// w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
	// w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
}
