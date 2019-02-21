package swagger

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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
