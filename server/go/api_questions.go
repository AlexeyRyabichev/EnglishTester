package swagger

import (
	"encoding/json"
	"log"
	"net/http"
)

func QuestionsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//TODO: more questions
	//email:=r.Header.Get("email")
	//check if this test is belong to this user
	var ques *[]Question
	err := db.Model((*Test)(nil)).Column("questions").Where("id = ?", 1).Select(&ques)
	if err != nil {
		log.Print(err)
	}
	jsoned, err := json.Marshal(ques)
	w.Write(jsoned)
	w.WriteHeader(http.StatusOK)
}

func AnswersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json charset=UTF=8")

	w.WriteHeader(http.StatusOK)
}
