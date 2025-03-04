package dto

type TransactionRequest struct {
	UserId                int
	TransactionDetailList []TransactionDetailRequest
}
