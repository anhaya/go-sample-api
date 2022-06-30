package test

import (
	"errors"
	"sync"
	"testing"

	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/anhaya/go-sample-api/internal/usecase/account"
	"github.com/anhaya/go-sample-api/internal/usecase/account/test/mocks"
	"github.com/stretchr/testify/suite"
)

type TestAccountCreateSuite struct {
	suite.Suite
	repository mocks.Repository
	useCase    account.AccountUseCase
}

func (suite *TestAccountCreateSuite) SetupTest() {
	suite.repository = mocks.Repository{}
	suite.useCase = account.NewAccount(&suite.repository)
}

func (suite *TestAccountCreateSuite) TestCreate_Success() {
	//given: prepare params
	documentNumber := "42311152434"
	balance := 100.00

	//and: mock repository
	suite.repository.On("Create", documentNumber, balance).Return(int64(1), nil)

	//when: call create usecase
	id, err := suite.useCase.Create(documentNumber, balance)

	//the: validate response
	suite.Nil(err)
	suite.repository.AssertNumberOfCalls(suite.T(), "Create", 1)
	suite.Equal(int64(1), id)
}

func (suite *TestAccountCreateSuite) TestCreate_GenerateAccountID_Success() {
	//given:

	//when: call generate account id
	accountId := account.GenerateAccountId()

	//then: validate is was generated
	suite.NotNil(accountId)
}

func (suite *TestAccountCreateSuite) TestCreate_GenerateAccountID_WithConcurrency() {
	//given:
	var accountIds []string
	var (
		wg      sync.WaitGroup
		workers = 100
		count   = 1000
		mutex   sync.Mutex
	)

	//when:
	// - Add multiple goroutines calling GenerateAccountId
	// - For each gorotuine, we call more 1000 times
	// - As they are concurrent. We have to sync the assign in accountsIds
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < count; j++ {
				accountId := account.GenerateAccountId()
				mutex.Lock()
				accountIds = append(accountIds, accountId)
				mutex.Unlock()
			}
		}()
	}
	// and: wait until all goroutines be finished
	wg.Wait()

	//then :
	// - check if we have some duplicated account id
	// - ps: It should have duplicated values in case we remove account id synchrounization
	hasDuplicatedAccountId := false
	mapDuplication := make(map[string]bool)

	for _, v := range accountIds {
		if _, ok := mapDuplication[v]; ok {
			hasDuplicatedAccountId = true
			break
		} else {
			mapDuplication[v] = true
		}
	}

	suite.False(hasDuplicatedAccountId)
}

func (suite *TestAccountCreateSuite) TestCreate_Error() {
	//given: prepare params
	documentNumber := "42311152434"
	balance := 100.00

	//and: mock repository
	mockRepository := errors.New("qualquer erro")
	suite.repository.On("Create", documentNumber, balance).Return(int64(0), mockRepository)

	//when: call create usecase
	id, err := suite.useCase.Create(documentNumber, balance)

	//the: validate response
	suite.Equal(err, entity.ErrInternalServer)
	suite.repository.AssertNumberOfCalls(suite.T(), "Create", 1)
	suite.Equal(int64(0), id)
}

func TestAccountCreate(t *testing.T) {
	suite.Run(t, new(TestAccountCreateSuite))
}
