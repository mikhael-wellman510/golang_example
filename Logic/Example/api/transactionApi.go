package api

import (
	"fmt"
	"latihan-golang/Logic/Example/controller"
	"latihan-golang/Logic/Example/dto"
)

func HitTransaction(transaction controller.TransactionController) {
	addTransaction := dto.TransactionRequest{
		UserId: 3,
		TransactionDetailList: []dto.TransactionDetailRequest{
			{ProductId: 2, Qty: 2},
			{ProductId: 3, Qty: 2},
			{ProductId: 4, Qty: 1},
		},
	}

	saveTransaction := transaction.CreateTransactionController(addTransaction)
	fmt.Println("Finish : ", saveTransaction)
}
