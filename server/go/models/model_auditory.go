package models

type Queue struct {
	Id int64
	Name string
	//тут пока что хз что но что-то обязательно будет
	//имя стуента
	//и хотелось ыб еще что-то полезное но пока хз
}

type Auditory struct {
	Id int64
	Number int
	Queue []Queue
	QueueId int64
}