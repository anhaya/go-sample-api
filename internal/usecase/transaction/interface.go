package transaction

import "github.com/anhaya/go-sample-api/internal/entity"

type UseCase interface {
	Create(accountId int, operationTypeId int, amount float64) error
}

type Repository interface {
	Create(transaction entity.Transaction) error
}
