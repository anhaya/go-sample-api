package db

import (
	"time"

	"github.com/anhaya/go-sample-api/internal/entity"
	db "github.com/anhaya/go-sample-api/pkg"
)

type transactionInfra struct {
	db db.DBExecutor
}

func NewTransaction(db db.DBExecutor) transactionInfra {
	return transactionInfra{
		db: db,
	}
}

func (t transactionInfra) Create(transaction entity.Transaction) error {
	_, err := t.db.Exec(
		"insert into transaction (account_id, operation_type_id, amount, event_date) values(?, ?, ?, ?)",
		transaction.AccountId,
		transaction.OperationTypeId,
		transaction.Amount,
		time.Now())

	if err != nil {
		return err
	}

	return nil
}
