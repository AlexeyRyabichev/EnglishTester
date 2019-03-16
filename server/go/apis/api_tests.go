package apis

import (
	"../DbWorker"
	Model "../models"
	"bufio"
	"encoding/json"
	"fmt"
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
		w.Write([]byte(err.Error()))
		return
	}

	_, err = DbWorker.Db.Model(&test).Insert()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TestDelete(w http.ResponseWriter, r *http.Request){
	//Todo: Run as transaction
	id:=r.Header.Get("testId")
	_, err := DbWorker.Db.Model((*Model.Test)(nil)).Where("id = ?",id).Delete()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	query:=	fmt.Sprintf("Select update_student_bytestId(%v)",id)
	_,err=DbWorker.Db.Exec(query)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}


func TestDocPost(w http.ResponseWriter,r *http.Request)  {
	fileQuestions,_,err := r.FormFile("questions")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
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
		return
	}
	answerBytes,err:=ioutil.ReadAll(fileAnswers)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	tmpfileQuestions, err := ioutil.TempFile("", "questions")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tmpfileAnswers, err := ioutil.TempFile("", "answers")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	defer os.Remove(tmpfileQuestions.Name()) // clean up
	defer os.Remove(tmpfileAnswers.Name())
	if _, err := tmpfileQuestions.Write(questionBytes); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if _, err := tmpfileAnswers.Write(answerBytes); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	test:= DocParser.GetTestFromDocx(tmpfileQuestions.Name(),tmpfileAnswers.Name())

	if err := tmpfileQuestions.Close(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err := tmpfileAnswers.Close(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_,err=DbWorker.Db.Model(test).Insert()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
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

func TestExportGet(w http.ResponseWriter, r *http.Request) {
	testId,err := strconv.ParseInt(mux.Vars(r)["testId"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}
	var test Model.Test
	err=DbWorker.Db.Model(&test).Where("id = ?",testId).Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}
	bytes,err:=DocParser.CreateTestDocx(&test)
	w.Header().Set("Content-Type","application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Write(bytes)

}
