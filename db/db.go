package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"tnwl_download/config"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.C.Db.Url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
