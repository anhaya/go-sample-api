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

func (a accountIfra) Create(documentNumber string, balance float64) (int64, error) {
	stmt, err := a.db.Prepare(`insert into account (document_number, balance) values(?,?)`)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(documentNumber, balance)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t accountIfra) Update(accountId int, newLimit float64) error {
	_, err := t.db.Exec("update account set balance = ? where id = ?", newLimit,
		accountId)

	if err != nil {
		return err
	}

	return nil
}

func (a accountIfra) Get(accountId int) (entity.Account, error) {
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
