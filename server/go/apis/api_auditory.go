package apis

import (
	"encoding/json"
	"log"
	"math"
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
	studentId := 157;
	var student Model.Student
	DbWorker.Db.Model(&student).Where("id = ?",studentId).Select()

	var auditories []Model.Auditory

	err:=DbWorker.Db.Model(&auditories).Select()
	if err != nil {

	}
	queueId:=findMinimalQueue(auditories)
	if(queueId==-1){
		log.Print("ERRORRRRRRRRR")
		return
	}

	var auditory Model.Auditory
	auditory.Id=queueId
	DbWorker.Db.Model(&auditory).WherePK().Select()
	if err != nil {

	}
	auditory.Queue=append(auditory.Queue, student)
	_,err=DbWorker.Db.Model(&auditory).Update()
	if err != nil {
		log.Print(err)
	}




}

func findMinimalQueue(auditories []Model.Auditory) int64 {
	var minId int64 = -1
	minLen:=math.MaxInt32
	for _,v:= range auditories{
		if(len(v.Queue)<=minLen){
			minId=v.Id;
			minLen=len(v.Queue)
		}
	}
	return minId
}

func AuditoriesGet(w http.ResponseWriter, r *http.Request)  {
	var aud []Model.Auditory
	err:=DbWorker.Db.Model(&aud).Select()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var prxAuds []Model.ProxyAuditory
	for _,v:=range aud{
		prxAud :=Model.ProxyAuditory{Id:v.Id}
		prxAud.Number= v.Number
		prxAud.Queue=make([]Model.Queue,len(v.Queue))
		for j:=0;j<len(v.Queue);j++  {
			prxAud.Queue[j]=Model.Queue{StudentId:v.Queue[j].Id,Name:v.Queue[j].Name}
		}
		prxAuds=append(prxAuds, prxAud)
	}

	jsoned,err :=json.Marshal(prxAuds)

	w.Write(jsoned)
	}


