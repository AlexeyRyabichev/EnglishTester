package apis

import (
	"../DbWorker"
	"../ExcelWorker"
	"../Roles"
	Model "../models"
	PassGenerator "../passGenerator"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func StudentsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	var students []Model.Student
	err := DbWorker.Db.Model(&students).Order("id ASC").Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	StudentJson, err := json.Marshal(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(StudentJson)
}

func StudentCreateWithExcelPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") != Roles.RolesText(Roles.Admin) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	file, fh, err := r.FormFile("file")
	if err != nil {
		log.Print(err)
	}
	size := fh.Size
	slice, err := ExcelWorker.ExcelAsSlice(file, size)
	var students []Model.Student
	for _, v := range slice[0][1:] {
		if(len(v)==0 || v ==nil){
			continue
		}
		var tmpStudent Model.Student
		tmpStudent.Name = v[0]
		tmpStudent.Email = v[1]
		pass, _ := PassGenerator.Generate(8, 8, 0, false, false)
		tmpStudent.Password = pass
		students = append(students, tmpStudent)
	}
	_, err = DbWorker.Db.Model(&students).Insert()
	if err != nil {
		log.Print(err)
	}
	w.WriteHeader(http.StatusOK)
}

func StudentPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") != Roles.RolesText(Roles.Admin) {
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
	if r.Header.Get("role") != Roles.RolesText(Roles.Admin) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	dec := json.NewDecoder(r.Body)
	var student Model.Student
	if err := dec.Decode(&student); err == io.EOF {
		//OK
	} else if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	_, err := DbWorker.Db.Model(&student).WherePK().Update()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func StudentsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") != Roles.RolesText(Roles.Admin) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	var students []Model.Student
	err := DbWorker.Db.Model(&students).Select()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	res, err := DbWorker.Db.Model(&students).Delete()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Println("deleted: ", res.RowsAffected())
	count, err := DbWorker.Db.Model((*Model.Student)(nil)).Count()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Println("left", count)
	w.WriteHeader(http.StatusOK)
}

func StudentDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") != Roles.RolesText(Roles.Admin) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	dec := json.NewDecoder(r.Body)
	var student Model.Student

	if err := dec.Decode(&student); err == io.EOF {
		//OK
	} else if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err := DbWorker.Db.Model(&student).WherePK().Delete()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func StudentsExportGet(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("role") == Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}

	var students []Model.Student

	DbWorker.Db.Model(&students).Select()

	excelFile := ExcelWorker.StudentsToExcel(students)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("content-disposition", "attachment; filename=StudentsList.xlsx")
	excelFile.Write(w)

	w.WriteHeader(http.StatusOK)
}

func StudentChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("role") != Roles.RolesText(Roles.Admin) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	studId := r.FormValue("id")
	var student Model.Student
	err := DbWorker.Db.Model(&student).Where("id = ?", studId).Select()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	pass, err := PassGenerator.Generate(8, 8, 0, false, false)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	student.Password = pass

	_, err = DbWorker.Db.Model(&student).WherePK().Update()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(pass))
	w.WriteHeader(http.StatusOK)
}
