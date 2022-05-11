package db

import (
	"database/sql"
	"fmt"
	"runtime"

	pkg "github.com/anhaya/go-sample-api/pkg"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	DB *sql.DB
}

func Open(user, password, host, port, database string) (*mysql, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return &mysql{}, err
	}
	return &mysql{
		DB: db,
	}, nil
}

func (s *mysql) Close() {
	s.DB.Close()
}

func (s *mysql) Atomic(fn func(dbexecutor pkg.DBExecutor) error) (err error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()

			switch e := p.(type) {
			case runtime.Error:
				panic(e)
			case error:
				err = fmt.Errorf("panic err: %v", p)
				return
			default:
				panic(e)
			}
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
