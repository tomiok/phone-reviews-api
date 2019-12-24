package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)

	if err != nil {
		_ = fmt.Errorf("cannot create db tentat: %s", err.Error())
		panic("..")
	}

	return db
}
