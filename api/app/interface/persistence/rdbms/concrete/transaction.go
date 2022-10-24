package concrete

import (
	"log"
)

//Rollback DbConn doesn't rollback, do nothing here
func (dbconn *DbConn) Rollback() error {
	return nil
}

//Commit DB doesnt commit, do nothing here
func (dbconn *DbConn) Commit() error {
	return nil
}

//TxEnd DB doesnt rollback, do nothing here
func (dbconn *DbConn) TxEnd(txFunc func() error) error {
	return nil
}

//TxEnd TX does rollback
func (txconn *TxConn) TxEnd(txFunc func() error) error {
	var err error
	tx := txconn.DB

	defer func() {
		if p := recover(); p != nil {
			log.Println("found p and rollback:", p)
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			log.Println("found error and rollback:", err)
			tx.Rollback() // err is non-nil; don't change it
		} else {
			log.Println("commit:")
			err = tx.Commit() // if Commit returns error update err with commit err
		}
	}()
	err = txFunc()
	return err
}

//Rollback TxConn rollback
func (txconn *TxConn) Rollback() error {
	return txconn.DB.Rollback()
}

//Commit TxConn rollback
func (txconn *TxConn) Commit() error {
	return txconn.DB.Commit()
}
