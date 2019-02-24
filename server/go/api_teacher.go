package swagger

import (
	"./Roles"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func TeachersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	var teachers []Teacher
	err := db.Model(&teachers).Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	teachersJson, err := json.Marshal(teachers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(teachersJson)

	w.WriteHeader(http.StatusOK)
}

func TeacherPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	dec := json.NewDecoder(r.Body)
	var teacher Teacher

	if err := dec.Decode(&teacher); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	_, err := db.Model(&teacher).Insert()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func TeacherDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}

	dec := json.NewDecoder(r.Body)
	var teacher Teacher

	if err := dec.Decode(&teacher); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	_, err := db.Model(&teacher).WherePK().Delete()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
