package dto

type OperationType string

var (
	OperationTypeIncome  OperationType = "income"
	OperationTypeExpense OperationType = "expense"
)

type AccountOperationDTO struct {
	id         int
	_type      OperationType
	accountId  int
	categoryId int
}
