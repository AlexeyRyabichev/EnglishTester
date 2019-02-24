package swagger

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var db *pg.DB

var passwordForDB = "tigra" //TODO: CHANGE THIS PASSWORD TO WHICH YOU HAVE ON YOUR SERVER

func InitDB() {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: passwordForDB,
	})
	err := createSchema(db)
	if err == nil {
		fmt.Print(err)
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
