package test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anhaya/go-sample-api/internal/app/handler"
	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/anhaya/go-sample-api/internal/usecase/account/test/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type TestAccountGetSuite struct {
	suite.Suite
	useCase mocks.UseCase
	handler handler.AccountHandler
}

func (suite *TestAccountGetSuite) SetupTest() {
	suite.useCase = mocks.UseCase{}
	suite.handler = handler.NewAccount(&suite.useCase)
}

func (suite *TestAccountGetSuite) TestGet_Succcess() {
	//given:

	//and: mock repository
	suite.useCase.On("Get", "1").Return(entity.Account{Id: "1"}, nil)

	//when: get account
	request, _ := http.NewRequest("GET", "/accounts", nil)
	vars := map[string]string{
		"accountId": "1",
	}
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	suite.handler.Get(response, request)

	//then: validate response
	suite.useCase.AssertNumberOfCalls(suite.T(), "Get", 1)
	suite.Equal(http.StatusOK, response.Code)
	//TODO - We should assert response body here as well
}

func (suite *TestAccountGetSuite) TestGet_ErrorUseCase() {
	//given:

	//and: mock usecase
	mockError := errors.New("core error")
	suite.useCase.On("Get", "1").Return(entity.Account{}, mockError)

	//when: get account
	request, _ := http.NewRequest("GET", "/accounts", nil)
	vars := map[string]string{
		"accountId": "1",
	}
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	suite.handler.Get(response, request)

	//then: validate response
	suite.useCase.AssertNumberOfCalls(suite.T(), "Get", 1)
	suite.Equal(http.StatusInternalServerError, response.Code)
	//TODO - We should assert response body here as well
}

func TestAccountGet(t *testing.T) {
	suite.Run(t, new(TestAccountGetSuite))
}
