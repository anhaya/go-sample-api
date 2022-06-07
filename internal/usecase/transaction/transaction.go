package transaction

import (
	"fmt"

	"github.com/anhaya/go-sample-api/internal/entity"
	infraDb "github.com/anhaya/go-sample-api/internal/infrastructure/db"
	account "github.com/anhaya/go-sample-api/internal/usecase/account"
	mysql "github.com/anhaya/go-sample-api/internal/usecase/mysql"
	pkgDb "github.com/anhaya/go-sample-api/pkg"
)

type TransactionUseCase struct {
	dbRepository   mysql.Repository
	tRepository    Repository
	accountUseCase account.UseCase
}

func NewTransaction(dbRepository mysql.Repository, tRepository Repository, accountUseCase account.UseCase) TransactionUseCase {
	return TransactionUseCase{
		dbRepository:   dbRepository,
		tRepository:    tRepository,
		accountUseCase: accountUseCase,
	}
}

func (t TransactionUseCase) Create(accountId int, operationTypeId int, amount float64) error {
	account, err := t.accountUseCase.Get(accountId)

	if err != nil {
		return entity.ErrInternalServer
	}

	var newBalance float64
	switch operationTypeId {
	case entity.SAQUE:
		newBalance = account.Balance - amount
	case entity.DEPOSITO:
		newBalance = account.Balance + amount
	default:
		return entity.ErrInternalServer
	}

	if newBalance < 0 {
		return entity.ErrInvalidLimit
	}

	err = t.dbRepository.Atomic(t.CreateAtomic(accountId, operationTypeId, amount, newBalance))
	if err != nil {
		return entity.ErrInternalServer
	}

	return nil
}

func (t TransactionUseCase) CreateAtomic(accountId int, operationTypeId int, amount float64, newBalance float64) func(dbexecutor pkgDb.DBExecutor) error {
	return func(dbexecutor pkgDb.DBExecutor) error {
		newTransactionRepo := infraDb.NewTransaction(dbexecutor)
		newAccountCore := infraDb.NewAccount(dbexecutor)

		transaction := entity.Transaction{
			AccountId:       accountId,
			OperationTypeId: operationTypeId,
			Amount:          amount,
		}
		err := newTransactionRepo.Create(transaction)

		if err != nil {
			fmt.Printf("error in persisting transaction `%s`", err)
			return entity.ErrInternalServer
		}

		err = newAccountCore.Update(transaction.AccountId, newBalance)
		if err != nil {
			fmt.Printf("error in updating transaction `%s`", err)
			return entity.ErrInternalServer
		}

		return err
	}
}
