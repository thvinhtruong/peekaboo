package registry

import (
	"database/sql"
	"server/app/interface/persistence/rdbms/concrete"
	"server/app/interface/persistence/rdbms/mysqldb"
	"server/app/usecase"
	"server/app/usecase/interactor"
)

func NewQuerier(NeedTransaction bool, db *sql.DB) mysqldb.Querier {
	conn := BuildTx(db, NeedTransaction)
	return mysqldb.Querier{DB: conn}
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

type UserAccessPoint struct {
	Service usecase.UserService
}

func BuildUserAccessPoint(NeedTransaction bool, db *sql.DB) *UserAccessPoint {
	querier := NewQuerier(NeedTransaction, db)
	usecaselayer := interactor.NewUserUsecase(&querier)

	return &UserAccessPoint{
		Service: usecaselayer,
	}
}

type ClassAccessPoint struct {
	Service usecase.ClassService
}

func BuildClassAccessPoint(NeedTransaction bool, db *sql.DB) *ClassAccessPoint {
	querier := NewQuerier(NeedTransaction, db)
	usecaselayer := interactor.NewClassUsecase(&querier)

	return &ClassAccessPoint{
		Service: usecaselayer,
	}
}

type TestAccessPoint struct {
	Service usecase.TestService
}

func BuildTestAccessPoint(NeedTransaction bool, db *sql.DB) *TestAccessPoint {
	querier := NewQuerier(NeedTransaction, db)
	usecaselayer := interactor.NewTestUsecase(&querier)

	return &TestAccessPoint{
		Service: usecaselayer,
	}
}

type TestResultAccessPoint struct {
	Service usecase.TestResultService
}

func BuildTestResultAccessPoint(NeedTransaction bool, db *sql.DB) *TestResultAccessPoint {
	querier := NewQuerier(NeedTransaction, db)
	usecaselayer := interactor.NewTestResultUsecase(&querier)

	return &TestResultAccessPoint{
		Service: usecaselayer,
	}
}

type AdminAccessPoint struct {
	Service usecase.AdminService
}

func BuildAdminAccessPoint(NeedTransaction bool, db *sql.DB) *AdminAccessPoint {
	querier := NewQuerier(NeedTransaction, db)
	usecaselayer := interactor.NewAdminUseCase(&querier)

	return &AdminAccessPoint{
		Service: usecaselayer,
	}
}
