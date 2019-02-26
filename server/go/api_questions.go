package swagger

import (
	"net/http"
)

func QuestionsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	////TODO: more questions
	//id := r.Header.Get("id")
	////email:=r.Header.Get("email")
	////check if this test is belong to this user
	//var student Student
	//var ques []Question
	//err := db.Model(&student).Relation("Test").Where("student.id = ?", id).Select()
	//if err != nil {
	//	log.Print(err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//ques = student.Test.Questions
	//jsoned, err := json.Marshal(ques)
	//w.Write(jsoned)
	w.WriteHeader(http.StatusOK)
}

func AnswersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json charset=UTF=8")

	w.WriteHeader(http.StatusOK)
}
