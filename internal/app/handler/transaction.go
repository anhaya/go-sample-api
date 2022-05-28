package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	customValidate "github.com/anhaya/go-sample-api/internal/app/validate"
	"github.com/anhaya/go-sample-api/internal/entity"
	"github.com/anhaya/go-sample-api/internal/usecase/transaction"
	"github.com/go-playground/validator/v10"
)

type TransactionHandler struct {
	usecase transaction.UseCase
}

func NewTransaction(usecase transaction.UseCase) TransactionHandler {
	return TransactionHandler{
		usecase: usecase,
	}
}

type CreateTransactionRequest struct {
	AccountId       int     `json:"account_id" validate:"required"`
	OperationTypeId int     `json:"operation_type_id" validate:"required,min=1,max=4"`
	Amount          float64 `json:"amount" validate:"required,NotNegative"`
}

// API Create Transaction godoc
// @Sumary Create Transaction
// @Description Create Transaction
// @Router /transactions [post]
// @Param Transaction body CreateTransactionRequest true "Transaction to create"
// @Accept json
// @Produce json
// @Success 201 "created"
// @Failure 400 {object} string
// @Failure 500 {object} string
func (t TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Printf("Error in parsing body `%s`", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(entity.ErrInternalServer.Error()))
		return
	}
	//validate request
	validate := validator.New()
	validate.RegisterValidation("NotNegative", customValidate.NotNegative)
	if err = validate.Struct(input); err != nil {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	//call create
	err = t.usecase.Create(input.AccountId, input.OperationTypeId, input.Amount)

	if err != nil {
		if err == entity.ErrInvalidLimit {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(entity.ErrInvalidLimit.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(entity.ErrInternalServer.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
