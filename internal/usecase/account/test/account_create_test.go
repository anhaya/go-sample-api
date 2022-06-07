package test

import (
	"errors"
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
