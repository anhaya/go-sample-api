package test

import (
	"errors"
	"testing"

	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/anhaya/go-sample-api/internal/usecase/account"
	"github.com/anhaya/go-sample-api/internal/usecase/account/test/mocks"
	"github.com/stretchr/testify/suite"
)

type TestAccountGetSuite struct {
	suite.Suite
	repository mocks.Repository
	usecase    account.AccountUseCase
}

func (suite *TestAccountGetSuite) SetupTest() {
	suite.repository = mocks.Repository{}
	suite.usecase = account.NewAccount(&suite.repository)
}

func (suite *TestAccountGetSuite) TestGet_Success() {
	//given: prepare params
	accountId := 1

	//and: mock repository
	suite.repository.On("Get", accountId).Return(entity.Account{Id: accountId}, nil)

	//when: call usecase
	account, err := suite.usecase.Get(accountId)

	//then: validate response
	suite.Nil(err)
	suite.Equal(accountId, account.Id)
	suite.repository.AssertNumberOfCalls(suite.T(), "Get", 1)
}

func (suite *TestAccountGetSuite) TestGet_ErrorNotFound() {
	//given: prepare params
	accountId := 1

	//and: mock repository
	suite.repository.On("Get", accountId).Return(entity.Account{}, nil)

	//when: call usecase
	account, err := suite.usecase.Get(accountId)

	//then: validate response
	suite.Equal(err, entity.ErrNotFoundAccount)
	suite.Empty(account)
	suite.repository.AssertNumberOfCalls(suite.T(), "Get", 1)
}

func (suite *TestAccountGetSuite) TestGet_ErrorRepository() {
	//given: prepare params
	accountId := 1

	//and: mock repository
	mockedError := errors.New("qualquer erro")
	suite.repository.On("Get", accountId).Return(entity.Account{}, mockedError)

	//when: call usecase
	account, err := suite.usecase.Get(accountId)

	//then: validate response
	suite.Equal(err, entity.ErrInternalServer)
	suite.Empty(account)
	suite.repository.AssertNumberOfCalls(suite.T(), "Get", 1)
}

func TestAccountGet(t *testing.T) {
	suite.Run(t, new(TestAccountGetSuite))
}
