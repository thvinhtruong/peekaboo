package mysqldb

import "server/app/interface/persistence/rdbms/concrete"

// Querier is the pointer to the repository.
type Querier struct {
	DB concrete.DBTX
}

// EnableTx allow the transaction to be initialized.
func (q *Querier) EnableTx(txFunc func() error) error {
	return q.DB.TxEnd(txFunc)
}

func NewQuerier(DB concrete.DBTX) *Querier {
	return &Querier{
		DB: DB,
	}
}
