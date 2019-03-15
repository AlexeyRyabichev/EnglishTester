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
		log.Print(err)
	}
	auditory.Number =auditoryNumber
	DbWorker.Db.Model(&auditory)
}