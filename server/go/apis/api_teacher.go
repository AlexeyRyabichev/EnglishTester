package apis

import (
	"../DbWorker"
	"../Roles"
	Model "../models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func TeachersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.RolesText(Roles.Student)  {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	var teachers []Model.Teacher
	err := DbWorker.Db.Model(&teachers).Order("id ASC").Select()
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
	if r.Header.Get("role") == Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	dec := json.NewDecoder(r.Body)
	var teacher Model.Teacher

	if err := dec.Decode(&teacher); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	_, err := DbWorker.Db.Model(&teacher).Insert()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func TeacherDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.RolesText(Roles.Student)  {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}

	dec := json.NewDecoder(r.Body)
	var teacher Model.Teacher

	if err := dec.Decode(&teacher); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	_, err := DbWorker.Db.Model(&teacher).WherePK().Delete()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func TeacherChangePassword(w http.ResponseWriter, r *http.Request)  {
	pass:=r.FormValue("password")
	id:=r.Header.Get("id")
	var teacher Model.Teacher

	err:=DbWorker.Db.Model(&teacher).Where("id = ?", id).Select()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	teacher.Password=pass
	_,err=DbWorker.Db.Model(&teacher).WherePK().Update()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}