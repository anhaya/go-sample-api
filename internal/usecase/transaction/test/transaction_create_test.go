package test

import (
	"errors"
	"testing"

	"github.com/anhaya/go-sample-api/internal/entity"
	mockAccount "github.com/anhaya/go-sample-api/internal/usecase/account/test/mocks"
	mockMysql "github.com/anhaya/go-sample-api/internal/usecase/mysql/test/mocks"
	"github.com/anhaya/go-sample-api/internal/usecase/transaction"
	mockTransaction "github.com/anhaya/go-sample-api/internal/usecase/transaction/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestTransactionCreateSuite struct {
	suite.Suite
	mysqlRepository       mockMysql.Repository
	transactionRepository mockTransaction.Repository
	accountUseCase        mockAccount.UseCase
	usecase               transaction.TransactionUseCase
}

func (suite *TestTransactionCreateSuite) SetupTest() {
	suite.mysqlRepository = mockMysql.Repository{}
	suite.transactionRepository = mockTransaction.Repository{}
	suite.accountUseCase = mockAccount.UseCase{}
	suite.usecase = transaction.NewTransaction(
		&suite.mysqlRepository,
		&suite.transactionRepository,
		&suite.accountUseCase)
}

func (suite *TestTransactionCreateSuite) TestCreate_SuccessDeposito() {
	//given: prepare params
	accountId := 1
	operationTypeId := entity.DEPOSITO
	amount := 100.00

	//and: mocks
	suite.accountUseCase.On("Get", accountId).Return(entity.Account{Id: 1, Balance: 200}, nil)
	suite.mysqlRepository.On("Atomic", mock.Anything).Return(nil)

	//when: call usecase
	err := suite.usecase.Create(accountId, operationTypeId, amount)

	//then: validate response
	suite.Nil(err)
	suite.accountUseCase.AssertNumberOfCalls(suite.T(), "Get", 1)
	suite.mysqlRepository.AssertNumberOfCalls(suite.T(), "Atomic", 1)
}

func (suite *TestTransactionCreateSuite) TestCreate_SuccessSaque() {
	//given: prepare params
	accountId := 1
	operationTypeId := entity.SAQUE
	amount := 100.00

	//and: mocks
	suite.accountUseCase.On("Get", accountId).Return(entity.Account{Id: 1, Balance: 200}, nil)
	suite.mysqlRepository.On("Atomic", mock.Anything).Return(nil)

	//when: call usecase
	err := suite.usecase.Create(accountId, operationTypeId, amount)

	//then: validate response
	suite.Nil(err)
	suite.accountUseCase.AssertNumberOfCalls(suite.T(), "Get", 1)
	suite.mysqlRepository.AssertNumberOfCalls(suite.T(), "Atomic", 1)
}

func (suite *TestTransactionCreateSuite) TestCreate_ErrorGetAccount() {
	//given: prepare params
	accountId := 1
	operationTypeId := entity.SAQUE
	amount := 100.00

	//and: mocks
	mockedError := errors.New("qualquer erro")
	suite.accountUseCase.On("Get", accountId).Return(entity.Account{}, mockedError)

	//when: call usecase
	err := suite.usecase.Create(accountId, operationTypeId, amount)

	//then: validate response
	suite.Equal(err, entity.ErrInternalServer)
	suite.accountUseCase.AssertNumberOfCalls(suite.T(), "Get", 1)
}

func (suite *TestTransactionCreateSuite) TestCreate_ErrorInvalidLimit() {
	//given: prepare params
	accountId := 1
	operationTypeId := entity.SAQUE
	amount := 100.00

	//and: mocks
	suite.accountUseCase.On("Get", accountId).Return(entity.Account{Id: 1, Balance: 50}, nil)

	//when: call usecase
	err := suite.usecase.Create(accountId, operationTypeId, amount)

	//then: validate response
	suite.Equal(err, entity.ErrInvalidLimit)
	suite.accountUseCase.AssertNumberOfCalls(suite.T(), "Get", 1)
}

func (suite *TestTransactionCreateSuite) TestCreate_ErrorAtomicRepository() {
	//given: prepare params
	accountId := 1
	operationTypeId := entity.SAQUE
	amount := 100

	//and: mocks
	mockedRepo := errors.New("qualquer erro")
	suite.accountUseCase.On("Get", accountId).Return(entity.Account{Id: 1, Balance: 200}, nil)
	suite.mysqlRepository.On("Atomic", mock.Anything).Return(mockedRepo)

	//when: call usecase
	err := suite.usecase.Create(accountId, operationTypeId, float64(amount))

	//then: validate response
	suite.Equal(err, entity.ErrInternalServer)
	suite.accountUseCase.AssertNumberOfCalls(suite.T(), "Get", 1)
	suite.mysqlRepository.AssertNumberOfCalls(suite.T(), "Atomic", 1)

}

func TestTransactionCreate(t *testing.T) {
	suite.Run(t, new(TestTransactionCreateSuite))
}
