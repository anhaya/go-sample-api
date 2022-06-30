package account

import "github.com/anhaya/go-sample-api/internal/entity"

type Repository interface {
	Create(accountID string, documentNumber string, balance float64) error
	Update(accountId string, balance float64) error
	Get(accountId string) (entity.Account, error)
}

type UseCase interface {
	Create(documentNumber string, balance float64) (string, error)
	Update(accountId string, balance float64) error
	Get(accountId string) (entity.Account, error)
}
