package swagger

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

var db *pg.DB

func InitDB() {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "tigra",
	})
	err := createSchema(db)
	if err == nil {
		fmt.Print("kk")
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Student)(nil), (*Test)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaTeachers() error {
	for _, model := range []interface{}{(*Teacher)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaTest() error {
	for _, model := range []interface{}{(*Test)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertTests() error {

	var tests []Test = []Test{
		Test{Id: 1,
			Questions: []Question{{1, "dsds"}, {1, "vopros"}, {2, "kekus"}},
			Answers:   []string{"otvet1", "otvet2", "otvet3"},
		},
	}
	jsoned, err := json.Marshal(tests)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(jsoned))
	_, err = db.Model(&tests).Insert()
	if err != nil {
		log.Print(err)
	}
	return err
}
