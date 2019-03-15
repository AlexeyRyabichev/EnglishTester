package apis

import (
	"../DbWorker"
	Model "../models"
	"../ExcelWorker"
	"encoding/json"
	"log"
	"net/http"
)

func ScoreGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := r.Header.Get("id")
	var student Model.Student
	var test Model.Test
	var studentAns Model.AnswerContainer
	err := DbWorker.Db.Model(&student).Relation("Test").Where("student.id = ?", id).
		Column("test.*").Select(&test)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = DbWorker.Db.Model(&student).Where("id = ?", id).Select()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	studentAns = student.Answers
	score, _ := CountScore(test.Answers, studentAns)
	DbWorker.Db.Model(&score).Insert()
	jsonedScore, err := json.Marshal(score)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonedScore)

}

func ScoresGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var res []Model.Result
	var students []Model.Student
	err := DbWorker.Db.Model(&students).Relation("Score").
		Column("student.id", "student.name").Select(&res)
	//studentsCount := len(students)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	jsonedRes, err := json.Marshal(res)
	log.Print(string(jsonedRes))
	w.Write(jsonedRes)

}

func ScoreExcelGet(w http.ResponseWriter, r *http.Request) {
	var res []Model.Result
	var students []Model.Student
	err := DbWorker.Db.Model(&students).Relation("Score").
		Column("student.id", "student.name").Select(&res)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	excelFile := ExcelWorker.ScoresToExcel(res)

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("content-disposition", "attachment; filename=Scores.xlsx")
	excelFile.Write(w)

	w.WriteHeader(http.StatusOK)
}

//TODO: что-то подумать о том как же все-таки(в какой момент проводить скоринг)

func CountScore(correctAnswers *Model.AnswerContainer, studAnswers Model.AnswerContainer) (score Model.Score, err error) {

	for i := range correctAnswers.Base {
		if correctAnswers.Base[i] == studAnswers.Base[i] {
			score.Base++
		}
	}
	score.BaseAmount = len(correctAnswers.Base)

	for i := range correctAnswers.Reading {
		if correctAnswers.Reading[i] == studAnswers.Reading[i] {
			score.Reading++
		}
	}
	score.ReadingAmount = len(correctAnswers.Reading)
	score.ListeningAmount=20
	score.WritingAmount=10
	score.Sum = score.Reading + score.Base
	score.SumAmount = score.BaseAmount+score.ListeningAmount+score.ReadingAmount+score.WritingAmount
	return score, nil
}
