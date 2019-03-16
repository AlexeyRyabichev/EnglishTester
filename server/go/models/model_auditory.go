package models


type Auditory struct {
	Id int64
	Number int
	Queue []Student
}

type ProxyAuditory struct {
	Id int64
	Number int
	Queue []Queue
}
type Queue struct{
	StudentId int64 `json:"id"`
	Name string `json:"name"`
}