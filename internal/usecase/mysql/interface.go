package mysql

import db "github.com/anhaya/go-sample-api/pkg"

type Repository interface {
	Atomic(fn func(dbexecutor db.DBExecutor) error) (err error)
}
