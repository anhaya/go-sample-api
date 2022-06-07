package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anhaya/go-sample-api/internal/app/handler"
	"github.com/anhaya/go-sample-api/internal/usecase/account/test/mocks"
	"github.com/stretchr/testify/suite"
)

type TestAccountCreateSuite struct {
	suite.Suite
	useCase mocks.UseCase
	handler handler.AccountHandler
}

func (suite *TestAccountCreateSuite) SetupTest() {
	suite.useCase = mocks.UseCase{}
	suite.handler = handler.NewAccount(&suite.useCase)
}

func (suite *TestAccountCreateSuite) TestCreate_Success() {
	//given: prepare params
	documentNumber := "45322536445"
	balance := 100.00

	//and: mock repository
	suite.useCase.On("Create", documentNumber, balance).Return(int64(1), nil)

	//when: create account
	body := buildAccountCreateInput(documentNumber, balance)
	request, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	suite.handler.Create(response, request)

	//then: validate response
	suite.useCase.AssertNumberOfCalls(suite.T(), "Create", 1)
	suite.Equal(http.StatusCreated, response.Code)
	//TODO - We should assert response body here as well

}

func (suite *TestAccountCreateSuite) TestCreate_InvalidParameter() {

	for _, v := range []struct {
		body            []byte
		responseCode    int
		responseMessage string
	}{
		{body: buildAccountCreateInput("", 100.00), responseCode: 400},
		{body: buildAccountCreateInput(" ", 100.00), responseCode: 400},
		{body: buildAccountCreateInput("1111111111111111", 100.00), responseCode: 400},
		{body: buildAccountCreateInput("41422443543", 0), responseCode: 400},
		{body: buildAccountCreateInput("41422443543", -100.00), responseCode: 400},
	} {
		//given: prepared params

		//when: create account
		request, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(v.body))
		response := httptest.NewRecorder()
		suite.handler.Create(response, request)

		//then: validate response
		suite.Equal(v.responseCode, response.Code)
		//TODO - We should assert response body here as well
	}
}

func (suite *TestAccountCreateSuite) TestCreate_ErrorUseCase() {
	//given: prepare params
	documentNumber := "45322536445"
	balance := 100.00

	mockError := errors.New("core error")

	//and: mock repository
	suite.useCase.On("Create", documentNumber, balance).Return(int64(0), mockError)

	//when: create account
	body := buildAccountCreateInput(documentNumber, balance)
	request, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	suite.handler.Create(response, request)

	//then: validate response
	suite.useCase.AssertNumberOfCalls(suite.T(), "Create", 1)
	suite.Equal(http.StatusInternalServerError, response.Code)
	//TODO - We should assert response body here as well

}

func buildAccountCreateInput(documentNumber string, balance float64) []byte {
	body, _ := json.Marshal(struct {
		DocumentNumber string  `json:"document_number"`
		Balance        float64 `json:"balance"`
	}{
		DocumentNumber: documentNumber,
		Balance:        balance,
	})
	return body
}

func TestAccountCreate(t *testing.T) {
	suite.Run(t, new(TestAccountCreateSuite))
}
