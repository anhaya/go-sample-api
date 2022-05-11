package account

import "github.com/anhaya/go-sample-api/internal/entity"

type Repository interface {
	Create(documentNumber string, balance float64) (int64, error)
	Update(accountId int, balance float64) error
	Get(accountId int) (entity.Account, error)
}

type UseCase interface {
	Create(documentNumber string, balance float64) (int64, error)
	Update(accountId int, balance float64) error
	Get(accountId int) (entity.Account, error)
}
