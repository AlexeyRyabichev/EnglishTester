package apis

import (
	"log"
	"net/http"
		"../DbWorker"
		Model "../models"
	"strconv"
)

func AuditoryPost(w http.ResponseWriter, r *http.Request) {
	numberStr:=r.Header.Get("number");
	var auditory Model.Auditory
	auditoryNumber,err:=strconv.Atoi(numberStr);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}
	auditory.Number =auditoryNumber
	var queue []Model.Student
	auditory.Queue = queue
	_,err=DbWorker.Db.Model(&auditory).Insert();
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AuditoryDelete(w http.ResponseWriter, r *http.Request) {
	auditoryNumber:=r.Header.Get("number");
	var auditory Model.Auditory
	_,err:=DbWorker.Db.Model(&auditory).Where("number = ?",auditoryNumber).Delete()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddToQueuePost(w http.ResponseWriter, r *http.Request)   {
	w.WriteHeader(http.StatusOK)
}