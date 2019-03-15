package apis

import (
	"../DbWorker"
	Roles "../Roles"
	Model "../models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func QuestionsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := r.Header.Get("id")
	//email:=r.Header.Get("email")
	//check if this test is belong to this user
	var student Model.Student
	var test Model.Test
	err := DbWorker.Db.Model(&student).Relation("Test").
		Where("student.id = ?", id).Column("test.*").Select(&test)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	questions := Model.ProxyQuestions{
		BaseQuestions:    test.BaseQuestions,
		ReadingQuestions: test.ReadingQuestions,
		Writing:          test.Writing,
	}

	//ques = student.Test.Questions
	jsoned, err := json.Marshal(questions)
	w.Write(jsoned)
	//w.WriteHeader(http.StatusOK)
}

func AnswersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json charset=UTF=8")
	if r.Header.Get("role") != Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("У вас нет полномочий для этого действия."))
		log.Println(err.Error())
		log.Print(http.StatusText(http.StatusForbidden))
		return
	}

	id := r.Header.Get("id")
	var student Model.Student
	var answers Model.AnswerContainer
	err:=DbWorker.Db.Model(&student).Relation("Test").Where("student.id = ?",id).Column("student.*","test.id").Select(&student)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&answers); err == io.EOF {
		//OK
	} else if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	student.Answers=answers

	_, err = DbWorker.Db.Model(&student).Column("answers").Where("id = ?", id).Update()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	newscore,err:=CountScore(student.Test.Answers,answers)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var score  Model.Score
	err = DbWorker.Db.Model(&student).Relation("Score").Where("student.id = ?", id).Column("score.id").Select(&score)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	newscore.Id=score.Id
	_, err = DbWorker.Db.Model(&newscore).WherePK().Update()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SendWritingGrade(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("role") == Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("У вас нет полномочий для этого действия."))
		if err != nil {
			log.Println(err.Error())
		}
		log.Print(http.StatusText(http.StatusForbidden))
		return
	}

	var id int
	var grade int
	id, _ = strconv.Atoi(r.Header.Get("studentid"))
	grade, _ = strconv.Atoi(r.Header.Get("grade"))
	var student Model.Student
	var score Model.Score
	err1 := DbWorker.Db.Model(&student).Relation("Score").Where("student.id = ?", id).Column("score.id").Select(&score)
	if err1 != nil {
		log.Print("Error1:")
		log.Print(err1)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err1.Error()))
	}

	if score.WritingAmount < grade {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Grade couldn't me more than MaximumGrade \n "+
			"Your:%v \t Maximum:%v", grade, score.WritingAmount)))
		log.Print(fmt.Sprintf("Grade couldn't me more than MaximumGrade \n "+
			"Your:%v \t Maximum:%v", grade, score.WritingAmount))
		return
	}

	score.Writing = grade

	_, err2 := DbWorker.Db.Model(&score).WherePK().Update()
	if err1 != nil {
		log.Print("Error2:")
		log.Print(err2)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err2.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

func SendListeningGrade(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("role") == Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("У вас нет полномочий для этого действия."))
		if err != nil {
			log.Println(err.Error())
		}
		log.Print(http.StatusText(http.StatusForbidden))
		return
	}

	var id int
	var grade int
	id, _ = strconv.Atoi(r.Header.Get("studentid"))
	grade, _ = strconv.Atoi(r.Header.Get("grade"))
	var student Model.Student
	var score Model.Score
	err1 := DbWorker.Db.Model(&student).Relation("Score").Where("student.id = ?", id).Column("score.id").Select(&score)
	if err1 != nil {
		log.Print("Error1:")
		log.Print(err1)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err1.Error()))
	}

	if score.ListeningAmount < grade {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Grade couldn't me more than MaximumGrade \n "+
			"Your:%v \t Maximum:%v", grade, score.ListeningAmount)))
		log.Print(fmt.Sprintf("Grade couldn't me more than MaximumGrade \n "+
			"Your:%v \t Maximum:%v", grade, score.ListeningAmount))
		return
	}
	score.Listening = grade
	_, err2 := DbWorker.Db.Model(&score).WherePK().Update()
	if err1 != nil {
		log.Print("Error2:")
		log.Print(err2)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err2.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
