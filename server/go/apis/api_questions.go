package apis

import (
	sw "../../go"
	Model "../models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func QuestionsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	////TODO: more questions
	id := r.Header.Get("id")
	//email:=r.Header.Get("email")
	//check if this test is belong to this user
	var student Model.Student
	var test Model.Test
	err := sw.Db.Model(&student).Relation("Test").
		Where("student.id = ?", id).Select(&test)
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
	w.WriteHeader(http.StatusOK)
}

func AnswersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json charset=UTF=8")
	id := r.Header.Get("id")
	var student Model.Student
	var answers Model.AnswerContainer

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&answers); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}
	_, err := sw.Db.Model(&student).Column("answers").Where("id = ?", id).Update()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
