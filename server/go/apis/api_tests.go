package apis

import (
	"../DbWorker"
	Model "../models"
	"bufio"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"../DocParser"
)

func TestPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//r.ParseForm()
	r.ParseForm()
	text := r.FormValue("testText")

	var test Model.Test
	err := json.Unmarshal([]byte(text), &test)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = DbWorker.Db.Model(&test).Insert()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TestDocPost(w http.ResponseWriter,r *http.Request)  {
	fileQuestions,_,err := r.FormFile("questions")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	fileAnswers,_,err :=r.FormFile("answers")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	questionBytes,err:=ioutil.ReadAll(fileQuestions)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	answerBytes,err:=ioutil.ReadAll(fileAnswers)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	tmpfileQuestions, err := ioutil.TempFile("", "questions")
	if err != nil {
		log.Fatal(err)
	}

	tmpfileAnswers, err := ioutil.TempFile("", "answers")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfileQuestions.Name()) // clean up
	defer os.Remove(tmpfileAnswers.Name())
	if _, err := tmpfileQuestions.Write(questionBytes); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfileAnswers.Write(answerBytes); err != nil {
		log.Fatal(err)
	}

	test:= DocParser.GetTestFromDocx(tmpfileQuestions.Name(),tmpfileAnswers.Name())

	if err := tmpfileQuestions.Close(); err != nil {
		log.Fatal(err)
	}
	if err := tmpfileAnswers.Close(); err != nil {
		log.Fatal(err)
	}

	_,err=DbWorker.Db.Model(test).Insert()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)

	}

func TestPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	text := r.FormValue("testText")

	var test Model.Test
	err := json.Unmarshal([]byte(text), &test)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = DbWorker.Db.Model(&test).WherePK().Update()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TestGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	testId, err := strconv.ParseInt(mux.Vars(r)["testId"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}
	var test Model.Test

	err = DbWorker.Db.Model(&test).Where("id = ?", testId).Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
	testsJson, err := json.Marshal(test)

	w.Write(testsJson)
	w.WriteHeader(http.StatusOK)
}

func TestsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var tests []Model.Test
	err := DbWorker.Db.Model(&tests).Order("id ASC").Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
	testsJson, err := json.Marshal(tests)
	w.Write(testsJson)
	w.WriteHeader(http.StatusOK)
}

func CheckCredentialsTeacherPost(w http.ResponseWriter, r *http.Request) {
	var teachers []Model.Teacher
	scanner := bufio.NewReader(r.Body)

	res, _, _ := scanner.ReadLine()
	login := string(res)

	res, _, _ = scanner.ReadLine()
	pass := string(res)

	err := DbWorker.Db.Model(&teachers).Where("login = ? and password = ?", login, pass).Select()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if len(teachers) == 0 {
		w.Write([]byte("no"))
	} else {
		w.Write([]byte("yes"))
	}
	w.WriteHeader(http.StatusOK)

}

func CheckCredentialsPost(w http.ResponseWriter, r *http.Request) {
	var student []Model.Student

	scanner := bufio.NewReader(r.Body)

	res, _, _ := scanner.ReadLine()
	login := string(res)

	res, _, _ = scanner.ReadLine()
	pass := string(res)

	err := DbWorker.Db.Model(&student).Where("email = ? and password = ?", login, pass).Select()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if len(student) == 0 {
		w.Write([]byte("no"))
	} else {
		w.Write([]byte("yes"))
	}
	w.WriteHeader(http.StatusOK)

}
