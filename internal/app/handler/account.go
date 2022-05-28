package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	customValidate "github.com/anhaya/go-sample-api/internal/app/validate"
	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/anhaya/go-sample-api/internal/usecase/account"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	core account.UseCase
}

func NewAccount(core account.UseCase) AccountHandler {
	return AccountHandler{
		core: core,
	}
}

type CreateAccountRequest struct {
	DocumentNumber string  `json:"document_number" validate:"required,NotEmpty,max=15"`
	Balance        float64 `json:"balance" validate:"required,NotNegative"`
}

type GetAccountResponse struct {
	ID             int     `json:"id"`
	DocumentNumber string  `json:"document_number"`
	Balance        float64 `json:"balance"`
}

// API Create Account godoc
// @Sumary Create Account
// @Description Create Account according to document number
// @Router /accounts [post]
// @Param Account body CreateAccountRequest true "Account to create"
// @Accept json
// @Produce json
// @Success 201 "created"
// @Failure 400 {object} string
// @Failure 500 {object} string
func (a AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Printf("Error in parsing body `%s`", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(entity.ErrInternalServer.Error()))
		return
	}
	//validate request
	validate := validator.New()
	validate.RegisterValidation("NotEmpty", customValidate.NotEmpty)
	validate.RegisterValidation("NotNegative", customValidate.NotNegative)
	if err = validate.Struct(input); err != nil {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	//call create
	id, err := a.core.Create(input.DocumentNumber, input.Balance)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(entity.ErrInternalServer.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	toJ := GetAccountResponse{
		ID: int(id),
	}

	if err := json.NewEncoder(w).Encode(toJ); err != nil {
		fmt.Printf("error in marshalling json `%s`", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(entity.ErrNotFoundAccount.Error()))
	}
}

// API Get Account godoc
// @Sumary Get Account
// @Description Get Account according to account id
// @Router /accounts/{accountId} [get]
// @Param accountId path string true "Account id to search"
// @Accept json
// @Produce json
// @Success 200 {object} GetAccountResponse
// @Failure 404 {object} string
// @Failure 500 {object} string
func (a AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	acountId, err := strconv.Atoi(params["accountId"])
	if err != nil {
		fmt.Printf("Error in getting accountId header as integer `%s`", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//call GET
	account, err := a.core.Get(acountId)

	if err != nil {
		if err == entity.ErrNotFoundAccount {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	toJ := GetAccountResponse{
		ID:             account.Id,
		DocumentNumber: account.DocumentNumber,
		Balance:        account.Balance,
	}
	if err := json.NewEncoder(w).Encode(toJ); err != nil {
		fmt.Printf("error in marshalling json `%s`", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(entity.ErrNotFoundAccount.Error()))
	}
}
