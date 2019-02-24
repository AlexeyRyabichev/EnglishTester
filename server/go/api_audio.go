package swagger

import (
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func AudioStudentIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/mpeg")
	if r.Header.Get("role") == "student" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}
	studId, err := strconv.ParseInt(mux.Vars(r)["studentId"], 10, 64)
	var path string

	err = db.Model((*Audio)(nil)).
		Column("path").
		Where("student_id = ?", studId).
		Select(&path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	audiofile, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(audiofile)
	w.WriteHeader(http.StatusOK)
}

func AudioStudentIdPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == "student" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("У вас нет полномочий для этого действия."))
		return
	}

	studId, err := strconv.ParseInt(mux.Vars(r)["studentId"], 10, 64)

	r.ParseMultipartForm(32 << 20)

	file, _, err := r.FormFile("file") //retrieve the file from form data
	defer file.Close()                 //close the file when we finish

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//this is path which  we want to store the file'
	err = os.MkdirAll("./audios/", os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fp, err := filepath.Abs("./audios/" + strconv.FormatInt(studId, 10) + ".mp3")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(fp)
	defer f.Close()
	io.Copy(f, file)
	var audio = Audio{
		StudentId: studId,
		Path:      fp,
	}
	_, err = db.Model(&audio).Insert()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
