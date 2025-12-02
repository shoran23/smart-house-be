package middleware

import (
	"errors"
	"fmt"
	"net/http"
)

func CheckSessionToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(w, "Login No Session Token Found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Session Token Error", http.StatusBadRequest)
			return
		}

		fmt.Println("Session Cookie: ", cookie.Value)

		next.ServeHTTP(w, r)
	})
}
