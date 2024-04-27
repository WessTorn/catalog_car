package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"os"
)

func ConnectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
	return db
}

//
//func CreateSchema(db *pg.DB) error {
//	for _, model := range []interface{}{(*Car)(nil), (*Owner)(nil)} {
//		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
//			IfNotExists: true,
//		})
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func CreateSchema(db *pg.DB) (err error) {
	err = db.Model(new(Car)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	err = db.Model(new(Owner)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	return
}
