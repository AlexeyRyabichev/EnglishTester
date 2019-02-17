package swagger

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var db *pg.DB

func InitDB() pg.DB {
	db = pg.Connect(&pg.Options{
		User:"postgres",
		Password:"tigra",
	})
	err := createSchema()
	if(err==nil){
		fmt.Print("")
	}
	return *db
}

func createSchema() error {
	for _, model := range []interface{}{(*Student)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}


func getAllStudents(db *pg.DB) (students []Student,err error) {
	err = db.Model(&students).Select()
	return students,err;

}