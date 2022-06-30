package db

import (
	entity "github.com/anhaya/go-sample-api/internal/entity"
	pkg "github.com/anhaya/go-sample-api/pkg"
)

type accountIfra struct {
	db pkg.DBExecutor
}

func NewAccount(db pkg.DBExecutor) accountIfra {
	return accountIfra{
		db: db,
	}
}

func (a accountIfra) Create(accountID string, documentNumber string, balance float64) error {
	stmt, err := a.db.Prepare(`insert into account (id, document_number, balance) values(?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(accountID, documentNumber, balance)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (t accountIfra) Update(accountId string, newBalance float64) error {
	_, err := t.db.Exec("update account set balance = ? where id = ?", newBalance,
		accountId)

	if err != nil {
		return err
	}

	return nil
}

func (a accountIfra) Get(accountId string) (entity.Account, error) {
	stmt, err := a.db.Prepare(`select id, document_number, balance from account where id = ?`)
	if err != nil {
		return entity.Account{}, err
	}
	var account entity.Account
	rows, err := stmt.Query(accountId)
	if err != nil {
		return entity.Account{}, err
	}
	for rows.Next() {
		rows.Scan(&account.Id, &account.DocumentNumber, &account.Balance)
	}
	return account, nil
}
