package contract

import "database/sql"

type DataManager interface {
	RepoManager
	Begin() (TransactionManager, error)
	Close() error
}

type TransactionManager interface {
	RepoManager
	Rollback() error
	Commit() error
	GetDbTransaction() *sql.Tx
}
