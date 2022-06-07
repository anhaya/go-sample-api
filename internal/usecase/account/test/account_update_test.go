package test

import (
	"errors"
	"testing"

	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/anhaya/go-sample-api/internal/usecase/account"
	"github.com/anhaya/go-sample-api/internal/usecase/account/test/mocks"
	"github.com/stretchr/testify/suite"
)

type TestAccountUpdateSuite struct {
	suite.Suite
	repository mocks.Repository
	useCase    account.AccountUseCase
}

func (suite *TestAccountCreateSuite) Setuptest() {
	suite.repository = mocks.Repository{}
	suite.useCase = account.NewAccount(&suite.repository)
}

func (suite *TestAccountCreateSuite) TestUpdate_Sucess() {
	//given: prepare params
	accountId := 1
	balance := 100.00

	//and: mock repository
	suite.repository.On("Update", accountId, balance).Return(nil)

	//when: cal update usecase
	err := suite.useCase.Update(accountId, balance)

	//then: validate response
	suite.Nil(err)
	suite.repository.AssertNumberOfCalls(suite.T(), "Update", 1)
}

func (suite *TestAccountCreateSuite) TestUpdate_Error() {
	//given: prepare params
	accountId := 1
	balance := 100.00

	//and: mock repository
	mockError := errors.New("qualquer erro")
	suite.repository.On("Update", accountId, balance).Return(mockError)

	//when: cal update usecase
	err := suite.useCase.Update(accountId, balance)

	//then: validate response
	suite.Equal(err, entity.ErrInternalServer)
	suite.repository.AssertNumberOfCalls(suite.T(), "Update", 1)

}

func TestAccountUpdate(t *testing.T) {
	suite.Run(t, new(TestAccountUpdateSuite))
}
