package swagger

import "github.com/go-pg/pg/orm"

func CreateAudioSchema() error {
	for _, model := range []interface{}{(*Audio)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil

}
