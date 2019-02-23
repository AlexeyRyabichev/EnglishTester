package swagger

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("password")

	var student Student
	var teacher Teacher
	err := db.Model(&student).Where("email = ? and password = ?", email, pass).Select()
	var role string = "student"
	if err != nil {
		err = db.Model(&teacher).Where("login = ? and password = ?", email, pass).Select()
		role = "teacher"
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Wrong login or password"))
			return
		}
	}
	token, err := getToken(email, role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating JWT token: " + err.Error()))
	}
	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token:" + token))

}

func authMiddleware(next http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if routeName == "LoginPost" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		email := claims.(jwt.MapClaims)["email"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("email", email)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}
