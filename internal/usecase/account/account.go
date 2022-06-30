package account

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/oklog/ulid"
)

var (
	mAccountID    sync.Mutex
	entropySource io.Reader
)

type AccountUseCase struct {
	accountRepository Repository
}

func init() {
	entropySource = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UTC().UnixNano())), 0)
}

func NewAccount(accountRepository Repository) AccountUseCase {
	return AccountUseCase{
		accountRepository: accountRepository,
	}
}

func (a AccountUseCase) Create(documentNumber string, balance float64) (string, error) {
	id := GenerateAccountId()
	err := a.accountRepository.Create(id, documentNumber, balance)
	if err != nil {
		fmt.Printf("error in persisting account `%s`", err)
		return "", entity.ErrInternalServer
	}
	return id, nil
}

func (a AccountUseCase) Update(accountId string, balance float64) error {
	err := a.accountRepository.Update(accountId, balance)
	if err != nil {
		fmt.Printf("error in updating account `%s`", err)
		return entity.ErrInternalServer
	}
	return nil
}

func (a AccountUseCase) Get(accountId string) (entity.Account, error) {
	account, err := a.accountRepository.Get(accountId)
	if err != nil {
		fmt.Printf("error in getting account `%s`", err)
		return entity.Account{}, entity.ErrInternalServer
	}

	if account.Id == "" {
		fmt.Print("none account has been found")
		return entity.Account{}, entity.ErrNotFoundAccount
	}
	return account, nil
}

func GenerateAccountId() (accountID string) {
	defer mAccountID.Unlock()
	mAccountID.Lock()
	accountID = ulid.MustNew(ulid.Timestamp(time.Now().UTC()), entropySource).String()
	return
}
