package mysqldb

import (
	"database/sql"
	"log"
	"os"
	"server/app/interface/persistence/rdbms/concrete"
	"server/utils/service"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func NewQuery(testname string) Querier {
	db, err := service.OpenConnection()
	if err != nil {
		log.Fatalf("fail with error: %v", err)
	}

	// the mysqldb queries default with no transactions.
	conn := BuildTx(db, false)
	return Querier{DB: conn}
}

// Flag true returns the Tx
func BuildTx(db *sql.DB, Flag bool) concrete.DBTX {
	if Flag {
		tx, err := db.Begin()
		if err != nil {
			return nil
		}
		return &concrete.TxConn{DB: tx}
	}

	return &concrete.DbConn{DB: db}
}
