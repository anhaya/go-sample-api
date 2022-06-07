package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anhaya/go-sample-api/internal/app/handler"
	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/anhaya/go-sample-api/internal/usecase/transaction/test/mocks"
	"github.com/stretchr/testify/suite"
)

type TestTransactionalCreateSuite struct {
	suite.Suite
	useCase mocks.UseCase
	handler handler.TransactionHandler
}

func (suite *TestTransactionalCreateSuite) SetupTest() {
	suite.useCase = mocks.UseCase{}
	suite.handler = handler.NewTransaction(&suite.useCase)
}

func (suite *TestTransactionalCreateSuite) TestCreate_Success() {
	//given: parameters
	body := buildTransactionalCreateInput(1, 1, 100)

	//and: mock usecase
	suite.useCase.On("Create", 1, 1, 100).Return(nil)

	//when: call create handler
	request, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	suite.handler.Create(response, request)

	//then: validate response
	suite.useCase.AssertNumberOfCalls(suite.T(), "Create", 1)
	suite.Equal(http.StatusCreated, response.Code)

}

func (suite *TestTransactionalCreateSuite) TestCreate_ErrorInvalidLimit() {
	//given: parameters
	body := buildTransactionalCreateInput(1, 1, 100)

	//and: mock usecase
	mockError := entity.ErrInvalidLimit
	suite.useCase.On("Create", 1, 1, 100).Return(mockError)

	//when: call create handler
	request, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	suite.handler.Create(response, request)

	//then: validate response
	suite.useCase.AssertNumberOfCalls(suite.T(), "Create", 1)
	suite.Equal(http.StatusBadRequest, response.Code)
	//TODO - We should assert response body here as well
}

func (suite *TestTransactionalCreateSuite) TestCreate_ErrorUseCase() {
	//given: parameters
	body := buildTransactionalCreateInput(1, 1, 100)

	//and: mock usecase
	mockError := errors.New("qualquer erro")
	suite.useCase.On("Create", 1, 1, 100).Return(mockError)

	//when: call create handler
	request, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	suite.handler.Create(response, request)

	//then: validate response
	suite.useCase.AssertNumberOfCalls(suite.T(), "Create", 1)
	suite.Equal(http.StatusInternalServerError, response.Code)
	//TODO - We should assert response body here as well
}

func (suite *TestTransactionalCreateSuite) TestCreate_InvalidParameter() {

	for _, v := range []struct {
		body            []byte
		responseCode    int
		responseMessage string
	}{
		{body: buildTransactionalCreateInput(0, 1, 100.00), responseCode: 400},
		{body: buildTransactionalCreateInput(1, 5, 100.00), responseCode: 400},
		{body: buildTransactionalCreateInput(1, 0, 100.00), responseCode: 400},
		{body: buildTransactionalCreateInput(1, -1, 100.00), responseCode: 400},
		{body: buildTransactionalCreateInput(1, 1, 0), responseCode: 400},
		{body: buildTransactionalCreateInput(1, 1, -100.00), responseCode: 400},
	} {
		//given: prepared params

		//when: create account
		request, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(v.body))
		response := httptest.NewRecorder()
		suite.handler.Create(response, request)

		//then: validate response
		suite.Equal(v.responseCode, response.Code)
		//TODO - We should assert response body here as well
	}
}

func buildTransactionalCreateInput(accountId int, operationTypeId int, amount float64) []byte {
	body, _ := json.Marshal(struct {
		AccountId       int     `json:"account_id"`
		OperationTypeId int     `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
	}{
		AccountId:       accountId,
		OperationTypeId: operationTypeId,
		Amount:          amount,
	})
	return body
}
func TestTransactionalCreate(t *testing.T) {
	suite.Run(t, new(TestAccountCreateSuite))
}
