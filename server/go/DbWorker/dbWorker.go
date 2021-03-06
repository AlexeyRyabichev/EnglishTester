package DbWorker

import (
	Model "../models"
	"encoding/json"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"math"
)

var Db *pg.DB

const dbPASS = "tigra"        //CHANGE HERE
const addr = "localhost:5432" //"138.68.78.205:5432"

//TODO override

func InitDB() {
	Db = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     "postgres",
		Password: dbPASS,
	})
}

func TokenExists(token string) (bool, error) {
	var student Model.Student
	var teacher Model.Teacher
	exists, err := Db.Model(&student).Where("access_token = ?", token).Exists()
	if exists {
		return exists, nil
	} else {
		exists, err = Db.Model(&teacher).Where("access_token = ?", token).Exists()
		if err != nil {
			log.Print(err)
		}
		return exists, nil
	}
	return false, nil
}

func GiveStudentToken(student *Model.Student, token string) error {
	student.AccessToken = token
	_, err := Db.Model(student).WherePK().Update()
	return err
}

func GiveTeacherToken(teacher *Model.Teacher, token string) error {
	teacher.AccessToken = token
	_, err := Db.Model(teacher).WherePK().Update()
	return err
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Model.Student)(nil), (*Model.Test)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaStudents() error {
	for _, model := range []interface{}{(*Model.Student)(nil)} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaTeachers() error {
	for _, model := range []interface{}{(*Model.Teacher)(nil)} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaTest() error {
	for _, model := range []interface{}{(*Model.Test)(nil)} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaAuditory() error {
	for _, model := range []interface{}{(*Model.Auditory)(nil)} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}


func CreateSchemaScore() error {
	for _, model := range []interface{}{(*Model.Score)(nil)} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaAudio() error {
	for _, model := range []interface{}{(*Model.Audio)(nil)} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func MockAnswers() error {
	var ans Model.AnswerContainer
	ans.Writing = "eto moe esse"
	ans.Reading = []Model.Answer{Model.Answer{Answer: "otvetchtenie", Id: 1},
		{Id: 2, Answer: "kek"}}
	ans.Base = []Model.Answer{Model.Answer{Answer: "asadas", Id: 1}}

	str, _ := json.Marshal(ans)
	log.Print(string(str))
	return nil
}

func AddToQueuePost()   {

	studentId := 164;
	var student Model.Student
	Db.Model(&student).Where("id = ?",studentId).Select()

	var auditories []Model.Auditory

	err:=Db.Model(&auditories).Select()
	if err != nil {

	}
	queueId:=findMinimalQueue(auditories)
	if(queueId==-1){
		log.Print("ERRORRRRRRRRR")
		return
	}

	var auditory Model.Auditory
	auditory.Id=queueId
	Db.Model(&auditory).WherePK().Select()
	if err != nil {

	}
	auditory.Queue=append(auditory.Queue, student)
	_,err=Db.Model(&auditory).WherePK().Update()
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


//
//func InsertTests() error {
//
//	var tests []Test = []Test{
//		Test{Id: 1,
//			Questions: []Question{{1, "dsds"}, {1, "vopros"}, {2, "kekus"}},
//			Answers:   []string{"otvet1", "otvet2", "otvet3"},
//		},
//		Test{Id: 2,
//			Questions: []Question{{1, "vopros1"}, {1, "vopros2"}, {2, "vopros3"}},
//			Answers:   []string{"otv1", "otv2", "otv3"},
//		},
//	}
//	jsoned, err := json.Marshal(tests)
//	if err != nil {
//		log.Print(err)
//	}
//	fmt.Println(string(jsoned))
//	_, err = db.Model(&tests).Insert()
//	if err != nil {
//		log.Print(err)
//	}
//	return err
//}
