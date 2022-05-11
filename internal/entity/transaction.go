package entity

const (
	DEBITO  = 1
	CREDITO = 2
)

type Transaction struct {
	Id              int
	AccountId       int
	OperationTypeId int
	Amount          float64
	EventDate       string
}
