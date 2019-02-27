package swagger

import (
	"./Roles"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stbyte,_:=ioutil.ReadAll(r.Body)
	log.Print(string(stbyte))
	email := r.FormValue("email")
	pass := r.FormValue("password")

	var student Student
	var teacher Teacher
	var id int64
	err := db.Model(&student).Where("email = ? and password = ?", email, pass).Select()
	var role Roles.Role = Roles.Student
	id = student.Id
	if err != nil {
		err = db.Model(&teacher).Where("email = ? and password = ?", email, pass).Select()
		role = Roles.Teacher
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Wrong login or password"))
			return
		}
	}
	token, err := getToken(email, role, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating JWT token: " + err.Error()))
		return
	}

	switch role {
	case Roles.Student:
		err = GiveStudentToken(&student, token)
	case Roles.Teacher:
			err = GiveTeacherToken(&teacher, token)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)

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
		id := int64(claims.(jwt.MapClaims)["id"].(float64))

		r.Header.Set("email", email)
		r.Header.Set("role", role)
		r.Header.Set("id", strconv.FormatInt(id, 10))

		next.ServeHTTP(w, r)
	})
}
