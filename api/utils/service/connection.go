package service

import (
	"database/sql"
	"fmt"
	"server/app/interface/persistence/rdbms/concrete"

	_ "github.com/go-sql-driver/mysql"
)

const (
	User     = "root"
	Password = "root"
	Host     = "127.0.0.1:3306"
	Name     = "ieltscenter"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Name,
	))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func BuildTx(db *sql.DB, Flag bool) concrete.DBTX {
	if Flag {
		tx, err := db.Begin()
		if err != nil {
			return nil
		}
		return &concrete.TxConn{DB: tx}
	} else {
		return &concrete.DbConn{DB: db}
	}

}
