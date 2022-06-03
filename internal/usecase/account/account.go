package account

import (
	"fmt"

	"github.com/anhaya/go-sample-api/internal/entity"
)

type AccountUseCase struct {
	accountRepository Repository
}

func NewAccount(accountRepository Repository) AccountUseCase {
	return AccountUseCase{
		accountRepository: accountRepository,
	}
}

func (a AccountUseCase) Create(documentNumber string, balance float64) (int64, error) {
	id, err := a.accountRepository.Create(documentNumber, balance)
	if err != nil {
		fmt.Printf("error in persisting account `%s`", err)
		return 0, entity.ErrInternalServer
	}
	return id, nil
}

func (a AccountUseCase) Update(accountId int, balance float64) error {
	err := a.accountRepository.Update(accountId, balance)
	if err != nil {
		fmt.Printf("error in updating account `%s`", err)
		return entity.ErrInternalServer
	}
	return nil
}

func (a AccountUseCase) Get(accountId int) (entity.Account, error) {
	account, err := a.accountRepository.Get(accountId)
	if err != nil {
		fmt.Printf("error in getting account `%s`", err)
		return entity.Account{}, entity.ErrInternalServer
	}

	if account.Id == 0 {
		fmt.Print("none account has been found")
		return entity.Account{}, entity.ErrNotFoundAccount
	}
	return account, nil
}
