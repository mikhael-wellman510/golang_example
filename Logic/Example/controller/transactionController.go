package controller

import (
	"latihan-golang/Logic/Example/dto"
	"latihan-golang/Logic/Example/service"
)

type TransactionController struct {
	TransactionDetailService *service.TransactionDetailService
}

func NewTransactionController(transactionService *service.TransactionDetailService) *TransactionController {

	return &TransactionController{
		TransactionDetailService: transactionService,
	}
}

func (tc *TransactionController) CreateTransactionController(trxRequest dto.TransactionRequest) dto.TransactionResponse {

	return tc.TransactionDetailService.AddTransactionDetail(trxRequest)
}
