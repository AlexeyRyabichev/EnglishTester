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

func StudentsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	var students []Model.Student
	err := DbWorker.Db.Model(&students).Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	StudentJson, err := json.Marshal(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(StudentJson)
	w.WriteHeader(http.StatusOK)
}

func StudentCreateWithArrayPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	dec := json.NewDecoder(r.Body)
	var stArr []Model.Student
	for {

		if err := dec.Decode(&stArr); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", stArr)
	}

	_, err := DbWorker.Db.Model(&stArr).Insert()
	log.Printf("dsds")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func StudentPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	dec := json.NewDecoder(r.Body)
	var student Model.Student

	if err := dec.Decode(&student); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	_, err := DbWorker.Db.Model(&student).Insert()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func StudentPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	dec := json.NewDecoder(r.Body)
	var student Model.Student
	if err := dec.Decode(&student); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	_, err := DbWorker.Db.Model(&student).WherePK().Update()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func StudentsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.Role(Roles.Student).String() {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	var students []Model.Student
	err := DbWorker.Db.Model(&students).Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	res, err := DbWorker.Db.Model(&students).Delete()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println("deleted: ", res.RowsAffected())
	count, err := DbWorker.Db.Model((*Model.Student)(nil)).Count()
	if err != nil {
		panic(err)
	}
	log.Println("left", count)
	w.WriteHeader(http.StatusOK)
}

func StudentDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dec := json.NewDecoder(r.Body)
	var student Model.Student

	if err := dec.Decode(&student); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	_, err := DbWorker.Db.Model(&student).WherePK().Delete()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
