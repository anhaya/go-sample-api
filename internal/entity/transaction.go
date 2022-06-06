package entity

const (
	SAQUE    = 1
	DEPOSITO = 2
)

type Transaction struct {
	Id              int
	AccountId       int
	OperationTypeId int
	Amount          float64
	EventDate       string
}
